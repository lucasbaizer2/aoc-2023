package util

import "encoding/json"

type HashSet[T comparable] struct {
	items map[T]struct{}
}

func NewHashSet[T comparable](capacity int) HashSet[T] {
	return HashSet[T]{
		items: make(map[T]struct{}, capacity),
	}
}

func NewHashSetFromSlice[T comparable](slice []T) HashSet[T] {
	set := HashSet[T]{
		items: make(map[T]struct{}, len(slice)),
	}
	for _, item := range slice {
		set.Add(item)
	}
	return set
}

func (set *HashSet[T]) Add(item T) {
	set.items[item] = struct{}{}
}

func (set *HashSet[T]) Remove(item T) {
	delete(set.items, item)
}

func (set *HashSet[T]) Contains(item T) bool {
	_, ok := set.items[item]
	return ok
}

func (set *HashSet[T]) Items() []T {
	keys := make([]T, 0, len(set.items))
	for key := range set.items {
		keys = append(keys, key)
	}
	return keys
}

func (set *HashSet[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(set.Items())
}
