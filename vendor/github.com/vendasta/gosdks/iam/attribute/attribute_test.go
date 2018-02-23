package attribute

import (
	"testing"
	"github.com/vendasta/gosdks/pb/iam/v1"
	"github.com/stretchr/testify/assert"
	"time"
	"github.com/golang/protobuf/ptypes"
)

func Test_Int(t *testing.T) {
	cases := []struct {
		name   string
		input  int64
		output *iam_v1.Attribute
	}{
		{
			name:  "Zero value creates correct proto",
			input: 0,
			output: &iam_v1.Attribute{
				Kind: &iam_v1.Attribute_IntAttribute{
					IntAttribute: 0,
				},
			},
		},
		{
			name:  "Non Zero value creates correct proto",
			input: 64,
			output: &iam_v1.Attribute{
				Kind: &iam_v1.Attribute_IntAttribute{
					IntAttribute: 64,
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := Int(c.input)
			assert.Equal(t, c.output, actual)
		})
	}
}

func Test_Float(t *testing.T) {
	cases := []struct {
		name   string
		input  float64
		output *iam_v1.Attribute
	}{
		{
			name:  "Zero value creates correct proto",
			input: 0,
			output: &iam_v1.Attribute{
				Kind: &iam_v1.Attribute_DoubleAttribute{
					DoubleAttribute: 0,
				},
			},
		},
		{
			name:  "Non Zero value creates correct proto",
			input: 64.4,
			output: &iam_v1.Attribute{
				Kind: &iam_v1.Attribute_DoubleAttribute{
					DoubleAttribute: 64.4,
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := Float(c.input)
			assert.Equal(t, c.output, actual)
		})
	}
}

func Test_String(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		output *iam_v1.Attribute
	}{
		{
			name:  "Zero value creates correct proto",
			input: "",
			output: &iam_v1.Attribute{
				Kind: &iam_v1.Attribute_StringAttribute{
					StringAttribute: "",
				},
			},
		},
		{
			name:  "Non Zero value creates correct proto",
			input: "test",
			output: &iam_v1.Attribute{
				Kind: &iam_v1.Attribute_StringAttribute{
					StringAttribute: "test",
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := String(c.input)
			assert.Equal(t, c.output, actual)
		})
	}
}

func Test_Bool(t *testing.T) {
	cases := []struct {
		name   string
		input  bool
		output *iam_v1.Attribute
	}{
		{
			name:  "Zero value creates correct proto",
			input: false,
			output: &iam_v1.Attribute{
				Kind: &iam_v1.Attribute_BoolAttribute{
					BoolAttribute: false,
				},
			},
		},
		{
			name:  "Non Zero value creates correct proto",
			input: true,
			output: &iam_v1.Attribute{
				Kind: &iam_v1.Attribute_BoolAttribute{
					BoolAttribute: true,
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := Bool(c.input)
			assert.Equal(t, c.output, actual)
		})
	}
}

func Test_Struct(t *testing.T) {
	cases := []struct {
		name   string
		input  map[string]*iam_v1.Attribute
		output *iam_v1.Attribute
	}{
		{
			name:  "Zero value creates correct proto",
			input: nil,
			output: &iam_v1.Attribute{
				Kind: &iam_v1.Attribute_StructAttribute{
					StructAttribute: &iam_v1.StructAttribute{
						Attributes: nil,
					},
				},
			},
		},
		{
			name: "Non Zero value creates correct proto",
			input: map[string]*iam_v1.Attribute{
				"test": Int(64),
			},
			output: &iam_v1.Attribute{
				Kind: &iam_v1.Attribute_StructAttribute{
					StructAttribute: &iam_v1.StructAttribute{
						Attributes: map[string]*iam_v1.Attribute{
							"test": Int(64),
						},
					},
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := Struct(c.input)
			assert.Equal(t, c.output, actual)
		})
	}
}

func Test_List(t *testing.T) {
	cases := []struct {
		name   string
		input  []*iam_v1.Attribute
		output *iam_v1.Attribute
	}{
		{
			name:  "Zero value creates correct proto",
			input: nil,
			output: &iam_v1.Attribute{
				Kind: &iam_v1.Attribute_ListAttribute{
					ListAttribute: &iam_v1.ListAttribute{
						Attributes: nil,
					},
				},
			},
		},
		{
			name: "Non Zero value creates correct proto",
			input: []*iam_v1.Attribute{
				Int(2),
				Bool(true),
			},
			output: &iam_v1.Attribute{
				Kind: &iam_v1.Attribute_ListAttribute{
					ListAttribute: &iam_v1.ListAttribute{
						Attributes: []*iam_v1.Attribute{
							Int(2),
							Bool(true),
						},
					},
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := List(c.input)
			assert.Equal(t, c.output, actual)
		})
	}
}

func Test_Timestamp(t *testing.T) {
	t1, _ := ptypes.TimestampProto(time.Time{})
	t2, _ := ptypes.TimestampProto(time.Date(2000, 06, 05, 5, 23, 30, 0, &time.Location{}))
	cases := []struct {
		name   string
		input  time.Time
		output *iam_v1.Attribute
	}{
		{
			name:  "Zero value creates correct proto",
			input: time.Time{},
			output: &iam_v1.Attribute{
				Kind: &iam_v1.Attribute_TimestampAttribute{
					TimestampAttribute: t1,
				},
			},
		},
		{
			name:   "Non Zero value creates correct proto",
			input:  time.Date(2000, 06, 05, 5, 23, 30, 0, &time.Location{}),
			output: &iam_v1.Attribute{
				Kind: &iam_v1.Attribute_TimestampAttribute{
					TimestampAttribute: t2,
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := Timestamp(c.input)
			assert.Equal(t, c.output, actual)
		})
	}
}

func Test_GeoPoint(t *testing.T) {
	type input struct {
		lat  float64
		long float64
	}

	cases := []struct {
		name string
		input input
		output *iam_v1.Attribute
	}{
		{
			name:  "Zero value creates correct proto",
			input: input{
				lat: 0,
				long: 0,
			},
			output: &iam_v1.Attribute{
				Kind: &iam_v1.Attribute_GeopointAttribute{
					GeopointAttribute: &iam_v1.GeoPointAttribute{
						Latitude: 0,
						Longitude: 0,
					},
				},
			},
		},
		{
			name: "Non Zero value creates correct proto",
			input: input{
				lat:  64,
				long: 64,
			},
			output: &iam_v1.Attribute{
				Kind: &iam_v1.Attribute_GeopointAttribute{
					GeopointAttribute: &iam_v1.GeoPointAttribute{
						Latitude: 64,
						Longitude: 64,
					},
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := GeoPoint(c.input.lat, c.input.long)
			assert.Equal(t, c.output, actual)
		})
	}
}
