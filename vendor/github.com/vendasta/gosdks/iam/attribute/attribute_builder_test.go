package attribute

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/vendasta/gosdks/pb/iam/v1"
	"time"
)

func Test_Builder(t *testing.T) {
	cases := []struct {
		name   string
		input  *Builder
		output *iam_v1.StructAttribute
	}{
		{
			name:  "Int adds to the builder",
			input: NewBuilder().Int("int", 64),
			output: &iam_v1.StructAttribute{
				Attributes: map[string]*iam_v1.Attribute{
					"int": Int(64),
				},
			},
		},
		{
			name:  "Double adds to the builder",
			input: NewBuilder().Float("double", 64.4),
			output: &iam_v1.StructAttribute{
				Attributes: map[string]*iam_v1.Attribute{
					"double": Float(64.4),
				},
			},
		},
		{
			name:  "String adds to the builder",
			input: NewBuilder().String("string", "test"),
			output: &iam_v1.StructAttribute{
				Attributes: map[string]*iam_v1.Attribute{
					"string": String("test"),
				},
			},
		},
		{
			name:  "Bool adds to the builder",
			input: NewBuilder().Bool("bool", true),
			output: &iam_v1.StructAttribute{
				Attributes: map[string]*iam_v1.Attribute{
					"bool": Bool(true),
				},
			},
		},
		{
			name:  "Ints adds a list to the builder",
			input: NewBuilder().Ints("listInts", []int64{5, 12, 53}),
			output: &iam_v1.StructAttribute{
				Attributes: map[string]*iam_v1.Attribute{
					"listInts": List([]*iam_v1.Attribute{Int(5), Int(12), Int(53)}),
				},
			},
		},
		{
			name:  "Floats adds a list to the builder",
			input: NewBuilder().Floats("listDoubles", []float64{5.5, 12.3, 53.2}),
			output: &iam_v1.StructAttribute{
				Attributes: map[string]*iam_v1.Attribute{
					"listDoubles": List([]*iam_v1.Attribute{Float(5.5), Float(12.3), Float(53.2)}),
				},
			},
		},
		{
			name:  "Strings adds a list to the builder",
			input: NewBuilder().Strings("listStrings", []string{"test1", "test2", "test3"}),
			output: &iam_v1.StructAttribute{
				Attributes: map[string]*iam_v1.Attribute{
					"listStrings": List([]*iam_v1.Attribute{String("test1"), String("test2"), String("test3")}),
				},
			},
		},
		{
			name:  "Bools adds a list to the builder",
			input: NewBuilder().Bools("listBools", []bool{true, false}),
			output: &iam_v1.StructAttribute{
				Attributes: map[string]*iam_v1.Attribute{
					"listBools": List([]*iam_v1.Attribute{Bool(true), Bool(false)}),
				},
			},
		},
		{
			name:  "List adds a list to the builder",
			input: NewBuilder().List("list", []*iam_v1.Attribute{Int(64), String("test")}),
			output: &iam_v1.StructAttribute{
				Attributes: map[string]*iam_v1.Attribute{
					"list": List([]*iam_v1.Attribute{Int(64), String("test")}),
				},
			},
		},
		{
			name:  "Struct adds a struct to the builder",
			input: NewBuilder().Struct("struct", NewBuilder().Int("innerInt", 64)),
			output: &iam_v1.StructAttribute{
				Attributes: map[string]*iam_v1.Attribute{
					"struct": Struct(map[string]*iam_v1.Attribute{"innerInt": Int(64)}),
				},
			},
		},
		{
			name:  "Time adds to the builder",
			input: NewBuilder().Time("time", time.Date(2000, 06, 05, 5, 23, 30, 0, &time.Location{})),
			output: &iam_v1.StructAttribute{
				Attributes: map[string]*iam_v1.Attribute{
					"time": Timestamp(time.Date(2000, 06, 05, 5, 23, 30, 0, &time.Location{})),
				},
			},
		},
		{
			name:  "GeoPoint adds to the builder",
			input: NewBuilder().GeoPoint("geopoint", 10, 15),
			output: &iam_v1.StructAttribute{
				Attributes: map[string]*iam_v1.Attribute{
					"geopoint": GeoPoint(10, 15),
				},
			},
		},
		{
			name: "Complex builder with everything is correctly created",
			input: NewBuilder().
				Int("int", 64).
				Float("double", 64.4).
				String("string", "test").
				Bool("bool", true).
				Ints("listInts", []int64{5, 12, 53}).
				Floats("listDoubles", []float64{5.5, 12.3, 53.2}).
				Strings("listStrings", []string{"test1", "test2", "test3"}).
				Bools("listBools", []bool{true, false}).
				List("list", []*iam_v1.Attribute{Int(64), String("test")}).
				Struct("struct", NewBuilder().Int("innerInt", 64)).
				Time("time", time.Date(2000, 06, 05, 5, 23, 30, 0, &time.Location{})).
				GeoPoint("geopoint", 10, 15),
			output: &iam_v1.StructAttribute{
				Attributes: map[string]*iam_v1.Attribute{
					"int":         Int(64),
					"double":      Float(64.4),
					"string":      String("test"),
					"bool":        Bool(true),
					"listInts":    List([]*iam_v1.Attribute{Int(5), Int(12), Int(53)}),
					"listDoubles": List([]*iam_v1.Attribute{Float(5.5), Float(12.3), Float(53.2)}),
					"listStrings": List([]*iam_v1.Attribute{String("test1"), String("test2"), String("test3")}),
					"listBools":   List([]*iam_v1.Attribute{Bool(true), Bool(false)}),
					"list":        List([]*iam_v1.Attribute{Int(64), String("test")}),
					"struct":      Struct(map[string]*iam_v1.Attribute{"innerInt": Int(64)}),
					"time":        Timestamp(time.Date(2000, 06, 05, 5, 23, 30, 0, &time.Location{})),
					"geopoint":    GeoPoint(10, 15),
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := c.input.Build()
			assert.Equal(t, c.output, actual)
		})
	}
}
