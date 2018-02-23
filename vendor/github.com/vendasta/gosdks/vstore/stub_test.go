package vstore

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"sort"
)

func Test_HappyPathStubIsFilterMatch(t *testing.T) {
	type i struct {
		lo *lookupOption
		keys []string
	}
	s := NewHappyPathStub()

    cases := []struct {
        name   string
        input  i
        output bool
    }{
        {
            name:  "Empty Filter matches",
            input: i{
				lo: &lookupOption{},
				keys: []string{"PID", "Market1", "AG-123"},
			},
			output: true,
        },
		{
			name:  "Mismatched Filter misses",
			input: i{
				lo: &lookupOption{filters: []string{"MOO"}},
				keys: []string{"PID", "Market1", "AG-123"},
			},
			output: false,
		},
		{
			name:  "Filter matches first key",
			input: i{
				lo: &lookupOption{filters: []string{"PID"}},
				keys: []string{"PID", "Market1", "AG-123"},
			},
			output: true,
		},
		{
			name:  "Filter matches nth key",
			input: i{
				lo: &lookupOption{filters: []string{"PID", "Market1"}},
				keys: []string{"PID", "Market1", "AG-123"},
			},
			output: true,
		},
		{
			name:  "Filter matches last key",
			input: i{
				lo: &lookupOption{filters: []string{"PID", "Market1", "AG-123"}},
				keys: []string{"PID", "Market1", "AG-123"},
			},
			output: true,
		},
		{
			name:  "Filter does not support partial by default",
			input: i{
				lo: &lookupOption{filters: []string{"P"}},
				keys: []string{"PID", "Market1", "AG-123"},
			},
			output: false,
		},
		{
			name:  "Partial Filter matches first key",
			input: i{
				lo: &lookupOption{filters: []string{"P"}, partialFilter: true},
				keys: []string{"PID", "Market1", "AG-123"},
			},
			output: true,
		},
		{
			name:  "Partial Filter matches nth key",
			input: i{
				lo: &lookupOption{filters: []string{"PID", "Market1", "AG-"}, partialFilter: true},
				keys: []string{"PID", "Market1", "AG-123"},
			},
			output: true,
		},
		{
			name:  "Mismatched Partial Filter misses first key",
			input: i{
				lo: &lookupOption{filters: []string{"A"}, partialFilter: true},
				keys: []string{"PID", "Market1", "AG-123"},
			},
			output: false,
		},
		{
			name:  "Mismatched Partial Filter misses nth key",
			input: i{
				lo: &lookupOption{filters: []string{"PID", "Market1", "NG-"}, partialFilter: true},
				keys: []string{"PID", "Market1", "AG-123"},
			},
			output: false,
		},
		{
			name:  "Range Filter matches",
			input: i{
				lo: &lookupOption{beginFilters: []string{"PID", "Market1", "AG-100"},
					endFilters: []string{"PID", "Market1", "AG-200"}, partialFilter: false},
				keys: []string{"PID", "Market1", "AG-123"},
			},
			output: true,
		},
		{
			name:  "Range Filter excludes before beginning",
			input: i{
				lo: &lookupOption{beginFilters: []string{"PID", "Market1", "AG-150"},
					endFilters: []string{"PID", "Market1", "AG-200"}, partialFilter: false},
				keys: []string{"PID", "Market1", "AG-123"},
			},
			output: false,
		},
		{
			name:  "Range Filter excludes after end",
			input: i{
				lo: &lookupOption{beginFilters: []string{"PID", "Market1", "AG-100"},
					endFilters: []string{"PID", "Market1", "AG-120"}, partialFilter: false},
				keys: []string{"PID", "Market1", "AG-123"},
			},
			output: false,
		},
		{
			name:  "Range Filter matches includes beginning",
			input: i{
				lo: &lookupOption{beginFilters: []string{"PID", "Market1", "AG-123"},
					endFilters: []string{"PID", "Market1", "AG-200"}, partialFilter: false},
				keys: []string{"PID", "Market1", "AG-123"},
			},
			output: true,
		},
		{
			name:  "Range Filter matches excludes end",
			input: i{
				lo: &lookupOption{beginFilters: []string{"PID", "Market1", "AG-100"},
					endFilters: []string{"PID", "Market1", "AG-123"}, partialFilter: false},
				keys: []string{"PID", "Market1", "AG-123"},
			},
			output: false,
		},
    }
    for _, c := range cases {
        t.Run(c.name, func(t *testing.T) {
            actual := s.isFilterMatch(c.input.lo, c.input.keys)
            assert.Equal(t, c.output, actual)
        })
    }
}

type model struct {}

func (m model) Schema() *Schema {
	return nil
}

func TestEntities_Sort(t *testing.T) {
	type i struct {
		e entities
	}

	eAB := &entity{model{}, &KeySet{"ns", "k", []string{"A", "B"}}}
	eAB3 := &entity{model{}, &KeySet{"ns", "k", []string{"A", "B", "3"}}}
	eAB4 := &entity{model{}, &KeySet{"ns", "k", []string{"A", "B", "4"}}}
	eAB5 := &entity{model{}, &KeySet{"ns", "k", []string{"A", "B", "5"}}}
	cases := []struct {
		name   string
		input  i
		output entities
	}{
		{
			name:  "Sorts same length",
			input: i{
				e: entities{eAB5, eAB4, eAB3},
			},
			output: entities{eAB3, eAB4, eAB5},
		},
		{
			name:  "Sorts different length 1",
			input: i{
				e: entities{eAB4, eAB, eAB3},
			},
			output: entities{eAB, eAB3, eAB4},
		},
		{
			name:  "Sorts different length 2",
			input: i{
				e: entities{eAB, eAB4, eAB3},
			},
			output: entities{eAB, eAB3, eAB4},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			sort.Sort(c.input.e)
			assert.Equal(t, c.output, c.input.e)
		})
	}

}
