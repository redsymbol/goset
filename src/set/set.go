package set

import (
	"fmt"
	"sort"
)

/*
Basic set type.

Rather than instantiating directly, create new set objects with NewSet.

Not safe for concurrent access, so you must wrap operations with a
lock if accessing from multiple goroutines.

*/
type Set struct {
	items map[interface{}]*struct{}
}

/*
Create a new Set object. Arguments ITEMS... will be the initial members of the set.

  Examples:
     emptySet := NewSet()
     intsSet  := NewSet(7, 12, 42)
     mixedSet := NewSet("foo", "bar", 42, 2.818)

*/
func NewSet(items ...interface{}) *Set {
	s := Set{
		make(map[interface{}]*struct{}),
	}
	for _, item := range items {
		s.Add(item)
	}
	return &s
}

/*
Check whether item is contained in the set.

*/
func (s *Set) Contains(item interface{}) bool {
	_, ok := s.items[item]
	return ok
}

/*
Add an item to the set.

If the item already exists inside the set, do nothing.

*/
func (s *Set) Add(item interface{}) {
	s.items[item] = nil
}

/*
Delete the item from the set, if it's in.

If the item is not in the set, do nothing.

See also Remove.

*/
func (s *Set) Discard(item interface{}) {
	delete(s.items, item)
}

/*
Delete the item from the set, if it's in.

If the item is not in the set, panic.

See also Discard.

*/
func (s *Set) Remove(item interface{}) {
	if !s.Contains(item) {
		panic(fmt.Sprintf("Set does not contain \"%v\"", item))
	}
	delete(s.items, item)
}

/*
Get the number of items contained in the set.

*/
func (s *Set) Len() int {
	return len(s.items)
}

/*
Create a new set containing those elements in both this AND another set.

*/
func (s *Set) Intersect(other *Set) *Set {
	newset := NewSet()
	for item, _ := range s.items {
		if other.Contains(item) {
			newset.Add(item)
		}
	}
	return newset
}

/*
Create a new set containing those elements in this set OR another set.

*/
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

/*
Create a slice consisting of the elements in the set.

Useful if you need to pass the set data to a function that requires a slice.

The order is undefined. See also Sorted.

*/
func (s *Set) Slice() []interface{} {
	slice := make([]interface{}, len(s.items))
	var ii int
	for item, _ := range s.items {
		slice[ii] = item
		ii += 1
	}
	return slice
}

/*
Create a slice consisting of elements in the set, sorted.

The set type allows mixed types, which raises questions about
comparision.  What Sorted does is render each element as a string, and
then sorts them lexicographically. If you need a different ordering -
for example, your set contains just integers, and you want them sorted
numerically - use Slice and then sort it directly.

*/
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

/*
True iff every member of this set is in other.
*/
func (s *Set) IsSubsetOf(other *Set) bool {
	for item, _ := range s.items {
		if ! other.Contains(item) {
			return false
		}
	}
	return true
}

/*
True iff every member of other set is in this set.
*/
func (s *Set) IsSupersetOf(other *Set) bool {
	for item, _ := range other.items {
		if ! s.Contains(item) {
			return false
		}
	}
	return true
}

/*
Create and return a shallow copy of this set.
*/
func (s *Set) Copy() *Set {
	clone := NewSet()
	for item, _ := range s.items {
		clone.Add(item)
	}
	return clone
}

/*
True iff this set has the same members as the other set.
*/
func (s *Set) Equals(other *Set) bool {
	if s.Len() == other.Len() {
		for item, _ := range s.items {
			if ! other.Contains(item) {
				return false
			}
		}
		return true
	}
	return false
}

/*
Return a new Set with elements in set that are not in the other.
*/
func (s *Set) Difference(other *Set) *Set {
	diff := NewSet()
	for item, _ := range s.items {
		if ! other.Contains(item) {
			diff.Add(item)
		}
	}
	return diff
}

/*
Remove and return an arbitrary element from the set.

Panic if set is empty.
*/
func (s *Set) Pop() interface{} {
	var item interface{}
	for item, _ = range s.items {
		break
	}
	delete(s.items, item)
	return item
}

/*
Remove all items from the set, making it empty.
*/
func (s *Set) Clear() {
	s.items = make(map[interface{}]*struct{})
}

/*
Return a new set that is the symmetric difference of this and the other.

In other words: Return a new set with elements in either the first OR
second set, but not both.

*/
func (s *Set) SymmetricDifference(other *Set) *Set {
	union := s.Union(other)
	intersect := s.Intersect(other)
	return union.Difference(intersect)
}