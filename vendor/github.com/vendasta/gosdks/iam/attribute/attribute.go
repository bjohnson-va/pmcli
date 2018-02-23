package attribute

import (
	"github.com/vendasta/gosdks/pb/iam/v1"
	"time"
	"github.com/golang/protobuf/ptypes"
)

// Int creates a new Int attribute
func Int(value int64) *iam_v1.Attribute {
	return &iam_v1.Attribute{
		Kind: &iam_v1.Attribute_IntAttribute{
			IntAttribute: value,
		},
	}
}

// Float creates a new Double attribute
func Float(value float64) *iam_v1.Attribute {
	return &iam_v1.Attribute{
		Kind: &iam_v1.Attribute_DoubleAttribute{
			DoubleAttribute: value,
		},
	}
}

// String creates a new String attribute
func String(value string) *iam_v1.Attribute {
	return &iam_v1.Attribute{
		Kind: &iam_v1.Attribute_StringAttribute{
			StringAttribute: value,
		},
	}
}

// Bool creates a new Bool attribute
func Bool(value bool) *iam_v1.Attribute {
	return &iam_v1.Attribute{
		Kind: &iam_v1.Attribute_BoolAttribute{
			BoolAttribute: value,
		},
	}
}

// Struct creates a new Struct attribute
func Struct(value map[string]*iam_v1.Attribute) *iam_v1.Attribute {
	return &iam_v1.Attribute{
		Kind: &iam_v1.Attribute_StructAttribute{
			StructAttribute: &iam_v1.StructAttribute{
				Attributes: value,
			},
		},
	}
}

// List creates a new List attribute
func List(values []*iam_v1.Attribute) *iam_v1.Attribute {
	return &iam_v1.Attribute{
		Kind: &iam_v1.Attribute_ListAttribute{
			ListAttribute: &iam_v1.ListAttribute{
				Attributes: values,
			},
		},
	}
}

// Timestamp creates a new Timestamp attribute
func Timestamp(value time.Time) *iam_v1.Attribute {
	t, _ := ptypes.TimestampProto(value)

	return &iam_v1.Attribute{
		Kind: &iam_v1.Attribute_TimestampAttribute{
			TimestampAttribute: t,
		},
	}
}

// GeoPoint creates a new GeoPoint attribute
func GeoPoint(lat, long float64) *iam_v1.Attribute {
	return &iam_v1.Attribute{
		Kind: &iam_v1.Attribute_GeopointAttribute{
			GeopointAttribute: &iam_v1.GeoPointAttribute{
				Latitude: lat,
				Longitude: long,
			},
		},
	}
}
