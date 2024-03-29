package set

import "github.com/cheekybits/genny/generic"

// Item the type of the Set
type Item generic.Type

// ItemSet the set of Items
type ItemSet struct {
	items map[Item]bool
}

// Add adds a new element to the Set. Returns a pointer to the Set.
func (s *ItemSet) Add(t Item) *ItemSet {
	if s.items == nil {
		s.items = make(map[Item]bool)
	}
	_, ok := s.items[t]
	if !ok {
		s.items[t] = true
	}
	return s
}

// Clear removes all elements from the Set
func (s *ItemSet) Clear() {
	s.items = make(map[Item]bool)
}

// Delete removes the Item from the Set and returns Has(Item)
func (s *ItemSet) Delete(item Item) bool {
	_, ok := s.items[item]
	if ok {
		delete(s.items, item)
	}
	return ok
}

// Has returns true if the Set contains the Item
func (s *ItemSet) Has(item Item) bool {
	_, ok := s.items[item]
	return ok
}

// Items returns the Item(s) stored
func (s *ItemSet) Items() []Item {
	items := []Item{}
	for i := range s.items {
		items = append(items, i)
	}
	return items
}

// Size returns the size of the set
func (s *ItemSet) Size() int {
	return len(s.items)
}

// Union operation
func (s *ItemSet) Union(s2 ItemSet) *ItemSet {
	s3 := ItemSet{}
	s3.items = make(map[Item]bool)

	for i := range s.items {
		s3.items[i] = true
	}

	for i := range s2.items {
		_, ok := s3.items[i]
		if !ok {
			s3.items[i] = true
		}
	}

	return &s3
}

func (s *ItemSet) Intersection(s2 ItemSet) *ItemSet {
	s3 := ItemSet{}
	s3.items = make(map[Item]bool)

	for i := range s2.items {
		_, ok := s.items[i]
		if ok {
			s3.items[i] = true
		}
	}

	return &s3
}

func (s *ItemSet) Difference(s2 ItemSet) *ItemSet {
	s3 := ItemSet{}
	s3.items = make(map[Item]bool)

	for i := range s.items {
		_, ok := s2.items[i]
		if !ok {
			s3.items[i] = true
		}
	}

	return &s3
}

func (s *ItemSet) Contains(s2 Item) bool {
	items := s.Items()
	for _, x := range items {
		if x == s2 {
			return true
		}
	}
	return false
}
