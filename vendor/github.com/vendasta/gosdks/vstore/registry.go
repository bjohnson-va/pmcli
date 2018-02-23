package vstore

import (
	"fmt"
	"reflect"
	"sync"
)

var registry = map[string]reflect.Type{}
var registryMu sync.Mutex

//RegisterModel associates a namespace and kind with a type that implements vstore.Model
//This should be called during every bootstrapping of an instance, as the association is not persisted
func RegisterModel(namespace string, kind string, m Model) {
	registryMu.Lock()
	defer registryMu.Unlock()
	registry[fmt.Sprintf("%s:%s", namespace, kind)] = reflect.TypeOf(m)
}

//ModelFrom returns the vstore.Model that is associated with a certain namespace and kind
func ModelFrom(namespace, kind string) Model {
	i := registry[fmt.Sprintf("%s:%s", namespace, kind)]; if i == nil {
		return nil
	}
	return reflect.New(i.Elem()).Interface().(Model)
}
