package set

import (
	"fmt"
	"sync"
)

type Set struct {
	items map[string]*struct{};
	mutex *sync.Mutex
}

func (s *Set) Contains(item string) bool {
	_, ok := s.items[item]
	return ok
}

func (s *Set) Add(item string) {
	s.items[item] = nil
}

/* Delete the item from the set, if it's in.
   If the item is not in the set, do nothing.
*/
func (s *Set) Discard(item string) {
	delete(s.items, item)
}

/* Delete the item from the set, if it's in.
   If the item is not in the set, panic.
*/
func (s *Set) Remove(item string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if ! s.Contains(item) {
		panic(fmt.Sprintf("Set does not contain \"%v\"", item))
	}
	delete(s.items, item)
}

func (s *Set) Len() int {
	return len(s.items)
}

func (s *Set) Intersect(other *Set) *Set {
	newset := NewSet()
	for item, _ := range s.items {
		if other.Contains(item) {
			newset.Add(item)
		}
	}
	return newset
}

func (s *Set) Union(other *Set) *Set {
	newset := NewSet()
	for item, _ := range s.items {
		newset.Add(item)
	}
	for item, _ := range other.items {
		newset.Add(item)
	}
	return newset
}

func NewSet(items ...string) *Set {
	s := Set{
		make(map[string]*struct{}),
		new(sync.Mutex),
	};
	for _, item := range items {
		s.Add(item)
	}
	return &s
}
