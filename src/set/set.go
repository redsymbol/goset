package set

type marker struct{}

type Set struct {
	items map[string]*marker;
}

func (s *Set) Contains(item string) bool {
	_, ok := s.items[item]
	return ok
}

func (s *Set) Add(item string) {
	s.items[item] = nil
}

func (s *Set) Delete(item string) {
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
	s := Set{make(map[string]*marker)};
	for _, item := range items {
		s.Add(item)
	}
	return &s
}
