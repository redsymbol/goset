package set

import (
	"fmt"
	"sync"
	"sort"
)

func NewSet(items ...interface{}) *Set {
	s := Set{
		make(map[interface{}]*struct{}),
		new(sync.Mutex),
	};
	for _, item := range items {
		s.Add(item)
	}
	return &s
}

type Set struct {
	items map[interface{}]*struct{};
	lock *sync.Mutex
}

func (s *Set) Contains(item interface{}) bool {
	_, ok := s.items[item]
	return ok
}

func (s *Set) Add(item interface{}) {
	s.items[item] = nil
}

/* Delete the item from the set, if it's in.
   If the item is not in the set, do nothing.
*/
func (s *Set) Discard(item interface{}) {
	delete(s.items, item)
}

/* Delete the item from the set, if it's in.
   If the item is not in the set, panic.
*/
func (s *Set) Remove(item interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
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

func (s *Set) Slice() []interface{} {
	slice := make([]interface{}, len(s.items))
	var ii int
	for item, _ := range s.items {
		slice[ii] = item
		ii += 1
	}
	return slice
}

func (s *Set) Sorted() []string {
	strslice := make([]string, len(s.items))
	for ii, val := range s.Slice() {
		switch val.(type) {
		case string:
			strslice[ii] = val.(string)
		default:
			strslice[ii] = fmt.Sprintf("%v", val)
		}
	}
	sort.Strings(strslice)
	return strslice
}