package attribute

import (
	"github.com/vendasta/gosdks/pb/iam/v1"
	"time"
)

// Builder
type Builder struct {
	attributes map[string]*iam_v1.Attribute
}

// Int add an integer value to the builder
func (ab *Builder) Int(key string, value int64) *Builder {
	ab.attributes[key] = Int(value)

	return ab
}

// Float add a double value to the builder
func (ab *Builder) Float(key string, value float64) *Builder {
	ab.attributes[key] = Float(value)

	return ab
}

// String add a string vlaue to the builder
func (ab *Builder) String(key string, value string) *Builder {
	ab.attributes[key] = String(value)

	return ab
}

// Bool add a bool value to the builder
func (ab *Builder) Bool(key string, value bool) *Builder {
	ab.attributes[key] = Bool(value)

	return ab
}

// Ints add a list of integers to the builder
func (ab *Builder) Ints(key string, values []int64) *Builder {
	attrs := make([]*iam_v1.Attribute, len(values))
	for i, v := range values {
		attrs[i] = Int(v)
	}
	return ab.List(key, attrs)
}

// Floats add a list of doubles to the builder
func (ab *Builder) Floats(key string, values []float64) *Builder {
	attrs := make([]*iam_v1.Attribute, len(values))
	for i, v := range values {
		attrs[i] = Float(v)
	}
	return ab.List(key, attrs)
}

// Strings add a list of strings to the builder
func (ab *Builder) Strings(key string, values []string) *Builder {
	attrs := make([]*iam_v1.Attribute, len(values))
	for i, v := range values {
		attrs[i] = String(v)
	}
	return ab.List(key, attrs)
}

// Bools add a list of bools to the builder
func (ab *Builder) Bools(key string, values []bool) *Builder {
	attrs := make([]*iam_v1.Attribute, len(values))
	for i, v := range values {
		attrs[i] = Bool(v)
	}
	return ab.List(key, attrs)
}

// List add a list to the builder
func (ab *Builder) List(key string, values []*iam_v1.Attribute) *Builder {
	ab.attributes[key] = List(values)
	return ab
}

// Struct add a struct to the builder
func (ab *Builder) Struct(key string, value *Builder) *Builder {
	ab.attributes[key] = &iam_v1.Attribute{
		Kind: &iam_v1.Attribute_StructAttribute{
			StructAttribute: value.Build(),
		},
	}
	return ab
}

// Time add time to the builder
func (ab *Builder) Time(key string, value time.Time) *Builder {
	ab.attributes[key] = Timestamp(value)

	return ab
}

// GeoPoint add a geo point to the builder
func (ab *Builder) GeoPoint(key string, lat float64, long float64) *Builder {
	ab.attributes[key] = GeoPoint(lat, long)

	return ab
}

// Build compile the builder into its proto
func (ab *Builder) Build() *iam_v1.StructAttribute {
	return &iam_v1.StructAttribute{
		Attributes: ab.attributes,
	}
}

// NewBuilder creates a new instance of the attribute builder
// Usage:
// attribute.NewBuilder().
//     String("UserID", "UID-123").
//     ListStrings("AccountGroupIDs", []string{"AG-123", "AG-234"}).
//     Struct("AccountGroupPermissions", attribute.NewBuilder().
//         Strings("AG-123", []string{"MP-123", "MP-ABC"}).
//         Strings("AG-234", []string{"MP-123", "MP-ABC"}),
//     ).
//     Build()
func NewBuilder() *Builder {
	return &Builder{
		attributes: map[string]*iam_v1.Attribute{},
	}
}
