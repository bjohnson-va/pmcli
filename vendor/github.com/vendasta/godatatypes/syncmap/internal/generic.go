package syncmap

import (
	"sync"
	"github.com/cheekybits/genny/generic"
)

// KeyType is the key type for a map.
type KeyType generic.Type

// ValueType is the value type for a map.
type ValueType generic.Type

// SyncKeyTypeValueTypeMap is a threadsafe wrapper around a map[KeyType]ValueType, safe for concurrent access.
type SyncKeyTypeValueTypeMap struct {
	// hashmap is the underlying map[KeyType]ValueType.
	hashmap map[KeyType]ValueType
	lock    sync.RWMutex
}

// NewSyncKeyTypeValueTypeMap creates a new threadsafe map storing ValueType values keyed by KeyType keys.
func NewSyncKeyTypeValueTypeMap() *SyncKeyTypeValueTypeMap {
	return ToSyncKeyTypeValueTypeMap(nil)
}

// SyncKeyTypeValueTypeMap creates a new SyncKeyTypeValueTypeMap prepopulated with the data from the
// specified map[KeyType]ValueType.
func ToSyncKeyTypeValueTypeMap(data map[KeyType]ValueType) *SyncKeyTypeValueTypeMap {
	if data == nil {
		data = make(map[KeyType]ValueType)
	}
	return &SyncKeyTypeValueTypeMap{hashmap: data}
}

// Get gets the ValueType value for the given KeyType key, and a bool indicating whether the key was present or not.
func (cm *SyncKeyTypeValueTypeMap) Get(k KeyType) (ValueType, bool) {
	cm.lock.RLock()
	v, ok := cm.hashmap[k]
	cm.lock.RUnlock()
	return v, ok
}

// Set sets the ValueType value for the specified KeyType key.
func (cm *SyncKeyTypeValueTypeMap) Set(k KeyType, v ValueType) {
	cm.lock.Lock()
	cm.hashmap[k] = v
	cm.lock.Unlock()
}

// Delete removes the specified KeyType key, returning its ValueType value and a bool indicating whether the delete
// occurred or not.
func (cm *SyncKeyTypeValueTypeMap) Delete(k KeyType) (ValueType, bool) {
	cm.lock.Lock()
	v, ok := cm.hashmap[k]
	if ok {
		delete(cm.hashmap, k)
	}
	cm.lock.Unlock()
	return v, ok
}
