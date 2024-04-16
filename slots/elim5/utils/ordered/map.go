package ordered

import (
	"golang.org/x/exp/slices"
)

// Map is an ordered map implemented using generics.
type Map[K comparable, V any] struct {
	keys   []K
	values map[K]V
}

// NewMap creates a new instance of Map.
func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		keys:   make([]K, 0),
		values: make(map[K]V),
	}
}

// Set adds or updates a key-value pair while maintaining the order of keys.
func (m *Map[K, V]) Set(key K, value V) {
	if _, ok := m.values[key]; !ok {
		m.keys = append(m.keys, key)
	}
	m.values[key] = value
}

// Get retrieves the value associated with the given key.
func (m *Map[K, V]) Get(key K) (V, bool) {
	val, ok := m.values[key]
	return val, ok
}

// DefaultGet retrieves the value for the given key, or returns the provided default value if the key does not exist.
func (m *Map[K, V]) DefaultGet(key K, defaultValue V) V {
	if val, ok := m.Get(key); ok {
		return val
	}
	return defaultValue
}

// Must retrieves the value associated with the given key or returns the zero value if the key does not exist.
func (m *Map[K, V]) Must(key K) V {
	return m.values[key]
}

// Exists checks if the specified key exists in the map.
func (m *Map[K, V]) Exists(key K) bool {
	_, ok := m.values[key]
	return ok
}

// Index returns the index of the specified key.
func (m *Map[K, V]) Index(key K) int {
	for i, k := range m.keys {
		if k == key {
			return i
		}
	}
	return -1
}

// GetByIndex retrieves the key-value pair at the specified index.
func (m *Map[K, V]) GetByIndex(index int) (K, V) {
	if index < 0 || index >= len(m.keys) {
		var zeroK K
		var zeroV V
		return zeroK, zeroV
	}
	key := m.keys[index]
	return key, m.values[key]
}

// Delete removes a key-value pair and the corresponding key from the order.
func (m *Map[K, V]) Delete(key K) {
	if _, ok := m.values[key]; ok {
		delete(m.values, key)
		for i, k := range m.keys {
			if k == key {
				m.keys = append(m.keys[:i], m.keys[i+1:]...)
				break
			}
		}
	}
}

// Len returns the number of key-value pairs in the map.
func (m *Map[K, V]) Len() int {
	return len(m.keys)
}

// Keys returns a slice containing all keys in the order they were added.
func (m *Map[K, V]) Keys() []K {
	return m.keys
}

// Values returns a slice containing all values in the order their keys were added.
func (m *Map[K, V]) Values() []V {
	var values []V
	for _, key := range m.keys {
		if val, ok := m.Get(key); ok {
			values = append(values, val)
		}
	}
	return values
}

// Clear removes all key-value pairs from the map.
func (m *Map[K, V]) Clear() {
	m.keys = []K{}
	m.values = make(map[K]V)
}

// Copy creates a shallow copy of the map.
func (m *Map[K, V]) Copy() *Map[K, V] {
	newMap := NewMap[K, V]()
	for _, key := range m.keys {
		newMap.Set(key, m.values[key])
	}
	return newMap
}

// Sort allows for external provision of a custom sort function.
func (m *Map[K, V]) Sort(sortFunc func(i, j K) bool) {
	temp := make([]K, len(m.keys))
	copy(temp, m.keys)
	slices.SortStableFunc(temp, sortFunc)
	m.keys = temp
}

// Range iterates over all key-value pairs, stopping if the callback returns false.
func (m *Map[K, V]) Range(fn func(K, V) bool) {
	for _, key := range m.keys {
		if !fn(key, m.values[key]) {
			break
		}
	}
}

// Swap swaps the positions of two keys if they exist.
func (m *Map[K, V]) Swap(iKey, jKey K) bool {
	return m.IndexSwap(m.Index(iKey), m.Index(jKey))
}

// IndexSwap swaps the positions of keys at the specified indices.
func (m *Map[K, V]) IndexSwap(i, j int) bool {
	if i < 0 || j < 0 || i >= len(m.keys) || j >= len(m.keys) {
		return false
	}
	m.keys[i], m.keys[j] = m.keys[j], m.keys[i]
	return true
}

// Insert inserts a key-value pair at the specified index.
func (m *Map[K, V]) Insert(i int, key K, value V) {
	if i < 0 || i > len(m.keys) {
		return
	}
	m.keys = append(m.keys[:i], append([]K{key}, m.keys[i:]...)...)
	m.values[key] = value
}

// Offset moves the key by the specified offset from its current position.
func (m *Map[K, V]) Offset(key K, offset int) {
	m.IndexOffset(m.Index(key), offset)
}

// IndexOffset moves the key at the specified index by the given offset.
func (m *Map[K, V]) IndexOffset(i int, offset int) {
	if i < 0 || i >= len(m.keys) {
		return
	}
	newIndex := i + offset
	if newIndex < 0 {
		newIndex = 0
	} else if newIndex >= len(m.keys) {
		newIndex = len(m.keys) - 1
	}
	m.indexMove(i, newIndex)
}

// Move moves a key to the specified position.
func (m *Map[K, V]) Move(key K, to int) bool {
	return m.IndexMove(m.Index(key), to)
}

// IndexMove moves the key at the specified index to a new position.
func (m *Map[K, V]) IndexMove(i, to int) bool {
	if i < 0 || to < 0 || i >= len(m.keys) || to >= len(m.keys) {
		return false
	}
	m.indexMove(i, to)
	return true
}

func (m *Map[K, V]) indexMove(i, to int) {
	keys := make([]K, len(m.keys))
	copy(keys, m.keys)
	key := keys[i]
	keys = append(keys[:i], keys[i+1:]...)
	keys = append(keys[:to], append([]K{key}, keys[to:]...)...)
	m.keys = keys
}

// GroupBy groups elements of a slice based on a provided key function.
// It creates a new Map where each key is generated by applying the key function
// to each element of the input slice, and the value is a slice of elements
// that share the same key. This allows for organizing elements into categories
// or groups efficiently.
func GroupBy[K comparable, V any, S []V](elements S, keySelector func(V) K) *Map[K, S] {
	m := NewMap[K, S]()
	for _, v := range elements {
		key := keySelector(v)
		// Append the current element to its corresponding group. Initialize a new group if necessary.
		m.Set(key, append(m.Must(key), v))
	}
	return m // Returns the Map containing the grouped elements.
}
