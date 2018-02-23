package mirage

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// CopyByPaths copies all fields specified in the list of paths from src to dest.
// See CopyByPath for more details.
func CopyByPaths(tagKey string, paths []string, src interface{}, dest interface{}) error {
	if src == nil || dest == nil {
		return errors.New("src and dest must be non-nil")
	}
	if reflect.TypeOf(src).Kind() != reflect.Ptr || reflect.TypeOf(dest).Kind() != reflect.Ptr {
		return fmt.Errorf("must pass pointers not struct for 'src' and 'dest', got: %s, %s", reflect.TypeOf(src).Kind(), reflect.TypeOf(dest).Kind())
	}

	splitUpPaths := splitPaths(paths)
	return copyByPath(tagKey, splitUpPaths, reflect.ValueOf(src), reflect.ValueOf(dest))
}

func splitPaths(paths []string) [][]string {
	results := make([][]string, len(paths))
	for i, path := range paths {
		results[i] = strings.Split(path, ".")
	}
	return results
}

func copyByPath(tagKey string, paths [][]string, src reflect.Value, dest reflect.Value) error {
	for _, path := range paths {
		err := copyByPathRecursive(tagKey, path, src, dest)
		if err != nil {
			return err
		}
	}
	return nil
}

func copyByPathRecursive(tagKey string, path []string, src reflect.Value, dest reflect.Value) error {
	// Get the actual structs that we're pointing to
	src, dest = reflect.Indirect(src), reflect.Indirect(dest)

	// This is a recursive function, if we've reached the end of the path we just copy the value and we're done.
	if len(path) == 0 {
		if src.Type() != dest.Type() {
			return fmt.Errorf("cannot copy differing types src: %s, dest: %s", src.Type(), dest.Type())
		}
		dest.Set(src)
		return nil
	}

	// pop one tag off the path
	tagValue := path[0]
	path = path[1:]

	// find the fields matching our tag in the struct
	srcField, err := getMatchingField(tagKey, tagValue, src)
	if err != nil {
		return err
	}
	destField, err := getMatchingField(tagKey, tagValue, dest)
	if err != nil {
		return err
	}
	// It's possible that the destination field is nil, but we probably still want to copy values inside.
	// If the field is a pointer type, and the destination is nil, then make it a pointer to an empty struct instead
	if destField.Kind() == reflect.Ptr && destField.IsNil() {
		destField.Set(reflect.New(destField.Type().Elem()))
	}

	// Check if the field we're going to recurse into is nil on the src side
	if srcField.Kind() == reflect.Ptr && srcField.IsNil() {
		// If we're at the end of the road we can just copy it over
		if len(path) == 0 {
			dest.Set(src)
			return nil
		}
		// We can't follow a nil path any further, return an error
		return fmt.Errorf("encountered nil along path in src at key \"%s\"", tagValue)
	}

	// follow the rest of the path
	err = copyByPathRecursive(tagKey, path, srcField, destField)
	if err != nil {
		return fmt.Errorf("%s > %s", tagValue, err.Error())
	}
	return nil
}

func getMatchingField(tagKey string, tagValue string, data reflect.Value) (reflect.Value, error) {
	typ := data.Type()
	for i := 0; i < typ.NumField(); i++ {
		currentTagValue := typ.Field(i).Tag.Get(tagKey)
		if currentTagValue == tagValue {
			return data.Field(i), nil
		}
	}
	return reflect.ValueOf(nil), fmt.Errorf("could not find field `%s:\"%s\" in type %s", tagKey, tagValue, typ.String())
}
