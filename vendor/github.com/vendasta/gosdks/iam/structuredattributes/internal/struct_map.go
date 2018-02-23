package internal

import (
	"sync"
	"reflect"
)

// structMap caches the model fields for a specific struct by the attribute tag identifier
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

func getFieldName(f reflect.StructField) string {
	alias, _ := f.Tag.Lookup("attribute")
	return alias
}

var smap = newStructMap()
