package mirage

// CopyByPath copies a field denoted by a tag from the src struct into the dest struct.
// For example, let's say we have the following struct with its vstore tag definitions:
//
//     type NestedStruct struct {
//       NestedString string `vstore:"my_nested_string"`
//     }
//
//     type AStruct struct {
//       StringField string         `vstore:"my_string_field"`
//       Nested      *NestedStruct  `vstore:nested"`
//     }
//
// CopyByPath takes a dot-notation path and will copy the value designated by the path from src to dest.
// It recurses into nested structs and will create new empty structs if the pointers are nil.
//
// Here's how we do it:
//
//     err := CopyByPath("vstore", "nested.my_nested_string", &src, &dest)
//
// This says to look inside the tag named `vstore` and to follow the fields matching vstore tags of `nested` and then
// inside that field look for `my_nested_string`. It ignores and DOES NOT overwrite any fields not specified in the field
// mask.
//
// NOTE: All fields and structs accessed must be PUBLIC (i.e. capitalized)
// NOTE: Source and Destination do NOT need to be the same type, so long as the path ends on the same types in each struct
func CopyByPath(tagKey string, path string, src interface{}, dest interface{}) error {
	return CopyByPaths(tagKey, []string{path}, src, dest)
}
