package internal

import (
	"github.com/vendasta/gosdks/iam/attribute"
	"reflect"
	"strconv"
)

// LegacyAttributeToModel sets all attributes from legacy onto to
func LegacyAttributeToModel(to interface{}, sa []*attribute.Attribute) error {
	sv := reflect.ValueOf(to)
	osv := reflect.ValueOf(to)
	if sv.Kind() == reflect.Ptr {
		osv = sv.Elem()
		if osv.Kind() != reflect.Struct {
			return newInvalidFieldError("expected struct got %v", osv.Kind())
		}
	}
	typ := osv.Type()
	m, err := smap.Get(typ)
	if err != nil {
		return err
	}
	for _, attr := range sa {
		f, ok := m[attr.Key]
		if !ok {
			continue
		}

		err := setLegacyValueOnField(f.Name, sv.Elem().FieldByName(f.Name), attr.Values)
		if err != nil {
			return err
		}
	}
	return nil
}

// setLegacyValueOnField sets a value on the interface specified
func setLegacyValueOnField(name string, field reflect.Value, values []string) error {
	if !field.IsValid() || !field.CanSet() {
		return newInvalidFieldError("Unable to use field %s, have you marked it private?", name)
	}
	v, err := getLegacyValueForField(field.Kind(), field, values)
	if err != nil {
		return err
	}

	field.Set(v.Convert(field.Type()))
	return nil
}

// getLegacyValueForField returns the value of a particular field
func getLegacyValueForField(k reflect.Kind, field reflect.Value, values []string) (*reflect.Value, error) {
	switch k {
	default:
		return nil, newInvalidFieldError("Type %s is not supported.", k.String())
	case reflect.Int, reflect.Int32, reflect.Int64:
		i, err := strconv.Atoi(values[0])
		if err != nil {
			return nil, err
		}
		val := reflect.ValueOf(i)
		return &val, nil
	case reflect.String:
		val := reflect.ValueOf(values[0])
		return &val, nil
	case reflect.Float64, reflect.Float32:
		i, err := strconv.ParseFloat(values[0], 64)
		if err != nil {
			return nil, err
		}
		val := reflect.ValueOf(i)
		return &val, nil
	case reflect.Bool:
		i, err := strconv.ParseBool(values[0])
		if err != nil {
			return nil, err
		}

		val := reflect.ValueOf(i)
		return &val, nil
	case reflect.Array, reflect.Slice:
		fieldType := reflect.TypeOf(field.Interface()).Elem()

		if fieldType.Kind() != reflect.String {
			return nil, newInvalidFieldError("List type %s is not supported.", fieldType.Kind().String())
		}
		val := reflect.ValueOf(values)
		return &val, nil
	}
}
