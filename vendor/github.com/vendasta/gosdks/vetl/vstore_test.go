package vetl

import (
	"github.com/vendasta/gosdks/pb/vetl/v1"
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/vendasta/gosdks/vstore"
)

func Test_GetPropertiesFromVStoreProperties(t *testing.T) {
	type test struct {
		name          string
		protos        []*vstore.Property
		expectedProps []*vetl_v1.Property
		expectedErr   error
	}

	cases := []*test{
		{
			name: "String vstore.Property returns string vetl_v1.Property",
			protos: []*vstore.Property{
				{
					Name: "prop",
					FType: vstore.StringType,
				},
			},
			expectedProps: []*vetl_v1.Property{
				{
					Name: "prop",
					Type: vetl_v1.Property_STRING,
				},
			},
		},
		{
			name: "Bool vstore.Property returns bool vetl_v1.Property",
			protos: []*vstore.Property{
				{
					Name: "prop",
					FType: vstore.BoolType,
				},
			},
			expectedProps: []*vetl_v1.Property{
				{
					Name: "prop",
					Type: vetl_v1.Property_BOOL,
				},
			},
		},
		{
			name: "Int64 vstore.Property returns int64 vetl_v1.Property",
			protos: []*vstore.Property{
				{
					Name: "prop",
					FType: vstore.IntType,
				},
			},
			expectedProps: []*vetl_v1.Property{
				{
					Name: "prop",
					Type: vetl_v1.Property_INT64,
				},
			},
		},
		{
			name: "Double vstore.Property returns double vetl_v1.Property",
			protos: []*vstore.Property{
				{
					Name: "prop",
					FType: vstore.FloatType,
				},
			},
			expectedProps: []*vetl_v1.Property{
				{
					Name: "prop",
					Type: vetl_v1.Property_DOUBLE,
				},
			},
		},
		{
			name: "Double vstore.Property returns double vetl_v1.Property",
			protos: []*vstore.Property{
				{
					Name: "prop",
					FType: vstore.TimeType,
				},
			},
			expectedProps: []*vetl_v1.Property{
				{
					Name: "prop",
					Type: vetl_v1.Property_TIMESTAMP,
				},
			},
		},
		{
			name: "Geopoint vstore.Property returns Geopoint vetl_v1.Property",
			protos: []*vstore.Property{
				{
					Name: "prop",
					FType: vstore.GeoPointType,
				},
			},
			expectedProps: []*vetl_v1.Property{
				{
					Name: "prop",
					Type: vetl_v1.Property_GEOPOINT,
				},
			},
		},
		{
			name: "Struct vstore.Property returns Struct vetl_v1.Property",
			protos: []*vstore.Property{
				{
					Name: "prop",
					FType: vstore.StructType,
				},
			},
			expectedProps: []*vetl_v1.Property{
				{
					Name: "prop",
					Type: vetl_v1.Property_STRUCT,
				},
			},
		},
		{
			name: "Struct vstore.Property with nested properties returns Struct vetl_v1.Property with nested properties",
			protos: []*vstore.Property{
				{
					Name: "prop",
					FType: vstore.StructType,
					Properties: []*vstore.Property{
						{
							Name:     "nested_prop",
							FType:     vstore.BoolType,
							IsRepeated: true,
						},
						{
							Name: "nested_geopoint",
							FType: vstore.GeoPointType,
						},
					},
				},
			},
			expectedProps: []*vetl_v1.Property{
				{
					Name: "prop",
					Type: vetl_v1.Property_STRUCT,
					Properties: []*vetl_v1.Property{
						{
							Name:     "nested_prop",
							Type:     vetl_v1.Property_BOOL,
							Repeated: true,
						},
						{
							Name: "nested_geopoint",
							Type: vetl_v1.Property_GEOPOINT,
						},
					},
				},
			},
		},
		{
			name: "Struct vstore.Property with nested structs returns Struct vetl_v1.Property with nested structs",
			protos: []*vstore.Property{
				{
					Name: "prop",
					FType: vstore.StructType,
					Properties: []*vstore.Property{
						{
							Name: "nested_struct",
							FType: vstore.StructType,
							Properties: []*vstore.Property{
								{
									Name:     "nested_prop",
									FType:     vstore.BoolType,
									IsRepeated: true,
								},
								{
									Name: "nested_geopoint",
									FType: vstore.GeoPointType,
								},
							},
						},
						{
							Name: "other_nested_struct",
							FType: vstore.StructType,
							Properties: []*vstore.Property{
								{
									Name: "nested_string_prop",
									FType: vstore.StringType,
								},
							},
						},
					},
				},
			},
			expectedProps: []*vetl_v1.Property{
				{
					Name: "prop",
					Type: vetl_v1.Property_STRUCT,
					Properties: []*vetl_v1.Property{
						{
							Name: "nested_struct",
							Type: vetl_v1.Property_STRUCT,
							Properties: []*vetl_v1.Property{
								{
									Name:     "nested_prop",
									Type:     vetl_v1.Property_BOOL,
									Repeated: true,
								},
								{
									Name: "nested_geopoint",
									Type: vetl_v1.Property_GEOPOINT,
								},
							},
						},
						{
							Name: "other_nested_struct",
							Type: vetl_v1.Property_STRUCT,
							Properties: []*vetl_v1.Property{
								{
									Name: "nested_string_prop",
									Type: vetl_v1.Property_STRING,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "vstore.Property with repeated value returns value with repeated value marked true",
			protos: []*vstore.Property{
				{
					Name: "prop",
					FType: vstore.StringType,
					IsRepeated: true,
				},
			},
			expectedProps: []*vetl_v1.Property{
				{
					Name: "prop",
					Type: vetl_v1.Property_STRING,
					Repeated: true,
				},
			},
		},
		{
			name: "vstore.Property with required value returns value with required value marked true",
			protos: []*vstore.Property{
				{
					Name: "prop",
					FType: vstore.StringType,
					IsRequired: true,
				},
			},
			expectedProps: []*vetl_v1.Property{
				{
					Name: "prop",
					Type: vetl_v1.Property_STRING,
					Required: true,
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			props, err := getPropertiesFromVStoreProperties(c.protos)
			assert.Equal(t, c.expectedErr, err)
			assert.Equal(t, c.expectedProps, props)
		})
	}
}