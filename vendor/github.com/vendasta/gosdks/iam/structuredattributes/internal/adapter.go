package internal

import (
	"github.com/vendasta/gosdks/pb/iam/v1"
	"reflect"
	"time"
	"github.com/golang/protobuf/ptypes"
)

// StructuredAttributeToModel sets all attributes from sa onto to
func StructuredAttributeToModel(to interface{}, sa *iam_v1.StructAttribute) error {
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
	for name, val := range sa.Attributes {
		f, ok := m[name]
		if !ok {
			continue
		}
		err := SetValueOnField(f.Name, sv.Elem().FieldByName(f.Name), val)
		if err != nil {
			return err
		}
	}
	return nil
}

// SetValueOnField sets a value on the interface specified
func SetValueOnField(name string, field reflect.Value, value *iam_v1.Attribute) error {
	if !field.IsValid() || !field.CanSet() {
		return newInvalidFieldError("Unable to use field %s, have you marked it private?", name)
	}
	v, err := GetValueForField(field.Kind(), field, value)
	if err != nil {
		return err
	}

	field.Set(v.Convert(field.Type()))
	return nil
}

//GetStructValue returns the value of a struct whose type is not known
func GetStructValue(field reflect.Value, sv interface{}, value *iam_v1.Attribute) (*reflect.Value, error) {
	switch svType := sv.(type) {
	default:
		t := reflect.TypeOf(sv)
		n := reflect.New(t)
		v, ok := value.Kind.(*iam_v1.Attribute_StructAttribute)
		if !ok {
			return nil, newInvalidFieldError("Expected struct but got %v. Value %v", svType, value)
		}
		err := StructuredAttributeToModel(n.Interface(), v.StructAttribute)
		if err != nil {
			return nil, err
		}
		return &n, nil
	case iam_v1.GeoPointAttribute:
		i, ok := value.Kind.(*iam_v1.Attribute_GeopointAttribute)
		if !ok {
			return nil, newInvalidFieldError("Expected geopoint but got %v.", value.Kind)
		}
		val := reflect.ValueOf(i.GeopointAttribute)
		return &val, nil
	case time.Time:
		i, ok := value.Kind.(*iam_v1.Attribute_TimestampAttribute)
		if !ok {
			return nil, newInvalidFieldError("Expected timestamp but got %v.", value.Kind)
		}
		t, err := ptypes.Timestamp(i.TimestampAttribute)
		if err != nil {
			return nil, err
		}
		if field.Kind() == reflect.Ptr {
			val := reflect.ValueOf(&t)
			return &val, nil
		}
		val := reflect.ValueOf(t)
		return &val, nil
	}
}

//GetValueForField returns the value of a particular field
func GetValueForField(k reflect.Kind, field reflect.Value, value *iam_v1.Attribute) (*reflect.Value, error) {
	switch k {
	default:
		return nil, newInvalidFieldError("Support for type %s is not supported.", k.String())
	case reflect.Int, reflect.Int32, reflect.Int64:
		i, ok := value.Kind.(*iam_v1.Attribute_IntAttribute)
		if !ok {
			return nil, newInvalidFieldError("Expected int but got %s.", value.Kind)
		}
		val := reflect.ValueOf(i.IntAttribute)
		return &val, nil
	case reflect.String:
		i, ok := value.Kind.(*iam_v1.Attribute_StringAttribute)
		if !ok {
			return nil, newInvalidFieldError("Expected string but got %s.", value.Kind)
		}
		val := reflect.ValueOf(i.StringAttribute)
		return &val, nil
	case reflect.Float64, reflect.Float32:
		i, ok := value.Kind.(*iam_v1.Attribute_DoubleAttribute)
		if !ok {
			return nil, newInvalidFieldError("Expected float but got %s.", value.Kind)
		}
		val := reflect.ValueOf(i.DoubleAttribute)
		return &val, nil
	case reflect.Bool:
		i, ok := value.Kind.(*iam_v1.Attribute_BoolAttribute)
		if !ok {
			return nil, newInvalidFieldError("Expected bool but got %s.", value.Kind)
		}
		val := reflect.ValueOf(i.BoolAttribute)
		return &val, nil
	case reflect.Struct:
		return GetStructValue(field, field.Interface(), value)
		case reflect.Array, reflect.Slice:
			fieldType := reflect.TypeOf(field.Interface()).Elem()

			lv, ok := value.Kind.(*iam_v1.Attribute_ListAttribute)
			if !ok {
				return nil, newInvalidFieldError("Expected list but got %s", value.Kind)
			}
			models := reflect.New(reflect.SliceOf(fieldType)).Interface()
			sl := reflect.ValueOf(models).Elem()

			for _, sv := range lv.ListAttribute.Attributes {
				if fieldType.Kind() == reflect.Struct {
					s, err := GetStructValue(field, reflect.Zero(fieldType).Interface(), sv)
					if err != nil {
						return nil, err
					}
					v := s.Convert(fieldType)
					sl.Set(reflect.Append(sl, v))
				} else if fieldType.Kind() != reflect.Ptr {
					lvi, err := GetValueForField(fieldType.Kind(), field, sv)
					if err != nil {
						return nil, err
					}
					v := lvi.Convert(fieldType)
					sl.Set(reflect.Append(sl, v))
				} else {
					using := reflect.Zero(reflect.TypeOf(field.Interface()).Elem())
					hm := reflect.New(using.Type()).Elem().Type().Elem()
					s, err := GetStructValue(field, reflect.Zero(hm).Interface(), sv)
					if err != nil {
						return nil, err
					}
					v := s.Convert(fieldType)
					sl.Set(reflect.Append(sl, v))
				}
			}
			return &sl, nil
		case reflect.Ptr:
			if field.IsNil() && field.Kind() != reflect.Slice {
				field.Set(reflect.New(field.Type().Elem()))
				using := field.Elem()
				return GetStructValue(field, using.Interface(), value)
			}
			return GetStructValue(field, field.Interface(), value)

	}
}
