package vstore

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/vendasta/gosdks/pb/vstorepb"
)

//ModelToByteArray accepts a instance of a struct that implements vstore.Model and returns a proto-marshaled byte array
//This allows you to encode an entity as a byte array or string.
func ModelToByteArray(m interface{}) ([]byte, error) {
	pb, err := ModelToStructPB(m)
	if err != nil {
		return nil, err
	}
	return Marshal(pb)
}

//ByteArrayToModel accepts a byte array and attempts to deserialize it into a vstore.Model according to the schema
//specified by the provided namespace-kind combination.
func ByteArrayToModel(namespace, kind string, bytes []byte) (Model, error) {
	pb := pool.Get().(*vstorepb.Struct)
	pb.Reset()
	defer pool.Put(pb)
	err := Unmarshal(bytes, pb)
	if err != nil {
		return nil, err
	}
	return StructPBToModel(namespace, kind, pb)
}

// ModelToStructPB transforms an arbitrary type implementing vstore.Model into a *vstorepb.Struct
func ModelToStructPB(m interface{}) (*vstorepb.Struct, error) {
	sv := reflect.ValueOf(m)
	osv := reflect.ValueOf(m)
	if sv.Kind() == reflect.Ptr {
		osv = sv.Elem()
		if osv.Kind() != reflect.Struct {
			return nil, fmt.Errorf("Expected struct got %v", osv.Kind())
		}
	}
	typ := osv.Type()

	mapCache, err := smap.Get(typ)
	if err != nil {
		return nil, err
	}
	s := &vstorepb.Struct{
		Values: make(map[string]*vstorepb.Value, len(mapCache)),
	}
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		alias := getFieldName(field)
		if alias == "" {
			continue
		}
		var val *vstorepb.Value
		var err error

		if sv.Kind() == reflect.Ptr || sv.Kind() == reflect.Interface {
			f := sv.Elem().FieldByName(field.Name)
			if !f.CanSet() {
				return nil, newInvalidFieldError("Cant serialize private field: `%s`.", field.Name)
			}
			val, err = ValueFromInterface(f.Interface())
			if err != nil {
				return nil, err
			}
		} else {
			val, err = ValueFromInterface(sv.FieldByName(field.Name).Interface())
			if err != nil {
				return nil, err
			}
		}
		if val == nil || val.GetKind() == nil {
			continue
		}
		s.Values[alias] = val
	}
	return s, nil
}

// StructPBToModel transforms a *vstorepb.Struct of a known namespace and kind into a vstore.Model
func StructPBToModel(namespace, kind string, s *vstorepb.Struct) (Model, error) {
	if s.Values == nil || len(s.Values) == 0 {
		return nil, nil
	}
	m := ModelFrom(namespace, kind)
	if m == nil {
		return nil, ErrUnregisteredModel
	}
	err := structPbToModel(m, s)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func getFieldName(f reflect.StructField) string {
	alias, ok := f.Tag.Lookup("vstore")
	if alias != "" {
		return alias
	}
	if alias == "" || !ok {
		alias, ok = f.Tag.Lookup("protobuf")
		if alias == "" || !ok {
			return ""
		}
		fields := strings.Split(alias, ",")
		for _, field := range fields {
			if strings.HasPrefix(field, "name=") {
				return field[5:]
			}
		}
	}
	return ""
}

func structPbToModel(to interface{}, s *vstorepb.Struct) error {
	sv := reflect.ValueOf(to)
	osv := reflect.ValueOf(to)
	if sv.Kind() == reflect.Ptr {
		osv = sv.Elem()
		if osv.Kind() != reflect.Struct {
			return fmt.Errorf("Expected struct got %v", osv.Kind())
		}
	}
	typ := osv.Type()
	m, err := smap.Get(typ)
	if err != nil {
		return err
	}
	for name, val := range s.Values {
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

//SetValueOnField sets a field to have a specified corresponding value
func SetValueOnField(name string, field reflect.Value, value *vstorepb.Value) error {
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

//GetValueForField returns the value of a particular field
func GetValueForField(k reflect.Kind, field reflect.Value, value *vstorepb.Value) (*reflect.Value, error) {
	switch k {
	default:
		return nil, newInvalidFieldError("Support for type %s is not supported.", k.String())
	case reflect.Int, reflect.Int32, reflect.Int64:
		i, ok := value.Kind.(*vstorepb.Value_IntValue)
		if !ok {
			return nil, newInvalidFieldError("Expected int but got %s from vStore.", value.Kind)
		}
		val := reflect.ValueOf(i.IntValue)
		return &val, nil
	case reflect.String:
		i, ok := value.Kind.(*vstorepb.Value_StringValue)
		if !ok {
			return nil, newInvalidFieldError("Expected string but got %s from vStore.", value.Kind)
		}
		val := reflect.ValueOf(i.StringValue)
		return &val, nil
	case reflect.Float64, reflect.Float32:
		i, ok := value.Kind.(*vstorepb.Value_DoubleValue)
		if !ok {
			return nil, newInvalidFieldError("Expected float but got %s from vStore.", value.Kind)
		}
		val := reflect.ValueOf(i.DoubleValue)
		return &val, nil
	case reflect.Bool:
		i, ok := value.Kind.(*vstorepb.Value_BoolValue)
		if !ok {
			return nil, newInvalidFieldError("Expected bool but got %s from vStore.", value.Kind)
		}
		val := reflect.ValueOf(i.BoolValue)
		return &val, nil
	case reflect.Struct:
		return GetStructValue(field, field.Interface(), value)
	case reflect.Array, reflect.Slice:
		fieldType := reflect.TypeOf(field.Interface()).Elem()

		// handle a byte array
		if fieldType.Kind() == reflect.Uint8 {
			i, ok := value.Kind.(*vstorepb.Value_BytesValue)
			if !ok {
				return nil, newInvalidFieldError("Expected []byte but got %s from vStore.", value.Kind)
			}
			val := reflect.ValueOf(i.BytesValue)
			return &val, nil
		}

		lv, ok := value.Kind.(*vstorepb.Value_ListValue)
		if !ok {
			return nil, newInvalidFieldError("Expected list but got %s from vStore.", value.Kind)
		}
		models := reflect.New(reflect.SliceOf(fieldType)).Interface()
		sl := reflect.ValueOf(models).Elem()

		for _, sv := range lv.ListValue.Values {
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

//GetStructValue returns the value of a struct whose type is not known
func GetStructValue(field reflect.Value, sv interface{}, value *vstorepb.Value) (*reflect.Value, error) {
	switch svType := sv.(type) {
	default:
		t := reflect.TypeOf(sv)
		n := reflect.New(t)
		v, ok := value.Kind.(*vstorepb.Value_StructValue)
		if !ok {
			return nil, newInvalidFieldError("Expected struct but got %v from vStore. Value %v", svType, value)
		}
		err := structPbToModel(n.Interface(), v.StructValue)
		if err != nil {
			return nil, err
		}
		return &n, nil
	case vstorepb.GeoPoint:
		i, ok := value.Kind.(*vstorepb.Value_GeopointValue)
		if !ok {
			return nil, newInvalidFieldError("Expected geopoint but got %v from vStore.", value.Kind)
		}
		val := reflect.ValueOf(i.GeopointValue)
		return &val, nil
	case time.Time:
		i, ok := value.Kind.(*vstorepb.Value_TimestampValue)
		if !ok {
			return nil, newInvalidFieldError("Expected timestamp but got %v from vStore.", value.Kind)
		}
		time, err := ptypes.Timestamp(i.TimestampValue)
		if err != nil {
			return nil, err
		}
		if field.Kind() == reflect.Ptr {
			val := reflect.ValueOf(&time)
			return &val, nil
		}
		val := reflect.ValueOf(time)
		return &val, nil
	}
}

//ValueFromInterface returns a vstorepb.Value from a generic value
func ValueFromInterface(i interface{}) (*vstorepb.Value, error) {
	switch iVal := i.(type) {
	default:
		typ := reflect.TypeOf(i)
		if typ.Kind() == reflect.Ptr {
			typ = typ.Elem()
		}
		if typ.Kind() == reflect.Struct {
			valOf := reflect.ValueOf(i)
			if valOf.Kind() == reflect.Ptr && valOf.IsNil() {
				return nil, nil
			}
			s, err := ModelToStructPB(i)
			if err != nil {
				return nil, err
			}

			return &vstorepb.Value{Kind: &vstorepb.Value_StructValue{StructValue: s}}, nil
		} else if typ.Kind() == reflect.Array || typ.Kind() == reflect.Slice {
			s := reflect.ValueOf(i)
			ret := make([]interface{}, s.Len())
			for i := 0; i < s.Len(); i++ {
				ret[i] = s.Index(i).Interface()
			}
			return ValueFromInterface(ret)
		}
		switch typ.Kind() {
		default:
			return nil, newInvalidFieldError("Unable to handle type %s for converstion to vstorepb.Value", typ.Kind().String())
		case reflect.String:
			return &vstorepb.Value{Kind: &vstorepb.Value_StringValue{StringValue: reflect.ValueOf(i).String()}}, nil
		case reflect.Int:
			return &vstorepb.Value{Kind: &vstorepb.Value_IntValue{IntValue: reflect.ValueOf(i).Int()}}, nil
		case reflect.Int32:
			return &vstorepb.Value{Kind: &vstorepb.Value_IntValue{IntValue: reflect.ValueOf(i).Int()}}, nil
		case reflect.Int64:
			return &vstorepb.Value{Kind: &vstorepb.Value_IntValue{IntValue: reflect.ValueOf(i).Int()}}, nil
		case reflect.Float32:
			return &vstorepb.Value{Kind: &vstorepb.Value_DoubleValue{DoubleValue: reflect.ValueOf(i).Float()}}, nil
		case reflect.Float64:
			return &vstorepb.Value{Kind: &vstorepb.Value_DoubleValue{DoubleValue: reflect.ValueOf(i).Float()}}, nil
		case reflect.Bool:
			return &vstorepb.Value{Kind: &vstorepb.Value_BoolValue{BoolValue: reflect.ValueOf(i).Bool()}}, nil
		}
	case *vstorepb.GeoPoint:
		if iVal == nil {
			return nil, nil
		}
		return &vstorepb.Value{Kind: &vstorepb.Value_GeopointValue{GeopointValue: iVal}}, nil
	case time.Time:
		if iVal.IsZero() {
			return nil, nil
		}
		tpb, err := ptypes.TimestampProto(iVal)
		if err != nil {
			return nil, err
		}
		return &vstorepb.Value{Kind: &vstorepb.Value_TimestampValue{TimestampValue: tpb}}, nil
	case *time.Time:
		if iVal == nil {
			return nil, nil
		}
		if iVal.IsZero() {
			return nil, nil
		}
		tpb, err := ptypes.TimestampProto(*iVal)
		if err != nil {
			return nil, err
		}
		return &vstorepb.Value{Kind: &vstorepb.Value_TimestampValue{TimestampValue: tpb}}, nil
	case []byte:
		return &vstorepb.Value{Kind: &vstorepb.Value_BytesValue{BytesValue: iVal}}, nil
	case []interface{}:
		lv := &vstorepb.ListValue{}
		if len(iVal) > 0 {
			lv.Values = make([]*vstorepb.Value, len(iVal))
		}
		for n, sub := range iVal {
			subV, err := ValueFromInterface(sub)
			if err != nil {
				return nil, err
			}
			lv.Values[n] = subV
		}
		return &vstorepb.Value{Kind: &vstorepb.Value_ListValue{
			ListValue: lv,
		}}, nil
	}
}

// structMap caches the model fields for a specific struct by the vstore tag identifier
type structMap struct {
	sync.Mutex
	maps map[reflect.Type]map[string]reflect.StructField
}

func (smap *structMap) Get(t reflect.Type) (map[string]reflect.StructField, error) {
	smap.Lock()
	defer smap.Unlock()
	return smap.get(t)
}

func (smap *structMap) get(t reflect.Type) (map[string]reflect.StructField, error) {
	mp, ok := smap.maps[t]
	if ok {
		return mp, nil
	}
	mp = map[string]reflect.StructField{}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		alias := getFieldName(field)
		if alias == "" {
			continue
		}
		mp[alias] = field
	}
	smap.maps[t] = mp
	return mp, nil
}

func newStructMap() *structMap {
	return &structMap{
		maps: map[reflect.Type]map[string]reflect.StructField{},
	}
}

var smap = newStructMap()

var pool = sync.Pool{
	New: func() interface{} {
		return &vstorepb.Struct{}
	},
}
