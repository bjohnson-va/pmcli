package attribute

// Deprecated: Use structured attributes instead
// Attribute represents a custom property on a subject
type Attribute struct {
	Key    string
	Values []string
}

// Deprecated: Use structured attributes instead
// NewLegacy returns a new Attribute
func NewLegacy(key string, values []string) *Attribute {
	return &Attribute{
		Key:    key,
		Values: values,
	}
}
