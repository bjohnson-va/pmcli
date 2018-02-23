package internal

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/vendasta/gosdks/iam/attribute"
)

func Test_LegacyAttributeToModel(t *testing.T) {
	type LegacyAttributeModel struct {
		Name            string                    `attribute:"name"`
		Version         int                       `attribute:"version"`
		Deleted         bool                      `attribute:"deleted"`
		Height          float64                   `attribute:"height"`
		AssociatedUsers []string                  `attribute:"associated_users"`
	}

	cases := []struct {
		name   string
		input  []*attribute.Attribute
		output LegacyAttributeModel
		error  error
	}{
		{
			name: "Extracts string attribute",
			input: []*attribute.Attribute{
				attribute.NewLegacy("name", []string{"Hank"}),
			},
			output: LegacyAttributeModel{
				Name: "Hank",
			},
		},
		{
			name: "Extracts int attribute",
			input: []*attribute.Attribute{
				attribute.NewLegacy("version", []string{"34"}),
			},
			output: LegacyAttributeModel{
				Version: 34,
			},
		},
		{
			name: "Extracts bool attribute",
			input: []*attribute.Attribute{
				attribute.NewLegacy("deleted", []string{"true"}),
			},
			output: LegacyAttributeModel{
				Deleted: true,
			},
		},
		{
			name: "Extracts float attribute",
			input: []*attribute.Attribute{
				attribute.NewLegacy("height", []string{"64.4"}),
			},
			output: LegacyAttributeModel{
				Height: 64.4,
			},
		},
		{
			name: "Extracts list attribute",
			input: []*attribute.Attribute{
				attribute.NewLegacy("associated_users", []string{"UID-123", "UID-343"}),
			},
			output: LegacyAttributeModel{
				AssociatedUsers: []string{"UID-123", "UID-343"},
			},
		},
		{
			name: "Extracts complex attribute",
			input: []*attribute.Attribute{
				attribute.NewLegacy("name", []string{"Hank"}),
				attribute.NewLegacy("version", []string{"34"}),
				attribute.NewLegacy("deleted", []string{"true"}),
				attribute.NewLegacy("height", []string{"64.4"}),
				attribute.NewLegacy("associated_users", []string{"UID-123", "UID-343"}),
			},
			output: LegacyAttributeModel{
				Name: "Hank",
				Version: 34,
				Deleted: true,
				Height: 64.4,
				AssociatedUsers: []string{"UID-123", "UID-343"},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := LegacyAttributeModel{}
			err := LegacyAttributeToModel(&actual, c.input)
			assert.Equal(t, c.error, err)
			assert.Equal(t, c.output, actual)
		})
	}
}
