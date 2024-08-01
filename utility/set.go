package utility

import (
	"sync"
)

type Set struct {
	sync.RWMutex
	items map[any]struct{}
}

func NewSet() *Set {
	return &Set{
		items: make(map[any]struct{}),
	}
}

// Add 向 Set 中加元素
func (s *Set) Add(element any) {
	s.Lock()
	defer s.Unlock()
	s.items[element] = struct{}{}
}

// Remove 從 Set 中移除元素
func (s *Set) Remove(element any) {
	s.Lock()
	defer s.Unlock()
	delete(s.items, element)
}

// Contains 檢查 Set 是否包含某個元素
func (s *Set) Contains(element any) bool {
	s.RLock()
	defer s.RUnlock()
	_, exists := s.items[element]
	return exists
}

// Clear 清空 Set 中所有元素
func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.items = make(map[any]struct{})
}

func (s *Set) Size() int {
	return len(s.items)
}

// ToSlice 回傳 Set 中所有元素
func (s *Set) ToSlice() []any {
	s.RLock()
	defer s.RUnlock()

	slice := make([]interface{}, 0, len(s.items))
	for item := range s.items {
		slice = append(slice, item)
	}

	return slice
}
