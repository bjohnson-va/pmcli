package util

import (
	"strings"
)

// FieldMask defines paths that will be updated in an update request
type FieldMask map[string]struct{}

// IsPropertyPresent Check if a property is present in the mask
func (f *FieldMask) IsPropertyPresent(property string) bool {
	return f.ArePropertiesPresent(property)
}

// ArePropertiesPresent Check if any property in the set of properties are present in the mask
func (f *FieldMask) ArePropertiesPresent(properties ...string) bool {
	if f == nil {
		// If a FieldMask is empty, it is supposed to mean that everything applies
		return true
	}
	for _, property := range properties {
		_, ok := (*f)[strings.ToLower(property)]
		if ok {
			return true
		}
	}
	return false
}

// AllPropertiesPresent Check if all properties in the set of properties are present in the mask
func (f *FieldMask) AllPropertiesPresent(properties ...string) bool {
	if f == nil {
		// If a FieldMask is empty, it is supposed to mean that everything applies
		return true
	}
	for _, property := range properties {
		_, ok := (*f)[strings.ToLower(property)]
		if !ok {
			return false
		}
	}
	return true
}

// FieldMaskProtoPathsToFieldMask convert field mask proto to field mask
func FieldMaskProtoPathsToFieldMask(paths []string) *FieldMask {
	if len(paths) == 0 {
		return nil
	}

	fm := make(FieldMask, len(paths))

	for _, path := range paths {
		if path != "" {
			// Adding support for "FieldMask", "fieldMask", "fieldmask", "field_mask", "Field_Mask" etc
			fm[strings.ToLower(ToSnakeCase(path))] = struct{}{}
			fm[strings.ToLower(ToCamelCase(path))] = struct{}{}
		}
	}

	if len(fm) == 0 {
		return nil
	}

	return &fm
}
