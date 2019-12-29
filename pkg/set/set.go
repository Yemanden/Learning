package set

import "sync"

// Set is interface, contains methods
// GetValue, Add, Remove, Union, Difference, Intersection, Subset
type Set interface {
	GetValues() map[int]bool
	Add(int)
	Remove(int) bool
	Union(Set) Set
	Difference(Set) Set
	Intersection(Set) Set
	Subset(Set) bool
}

type set struct {
	sync.RWMutex
	data map[int]bool
}

// GetValues returns values of the data contained in this Set
func (s *set) GetValues() map[int]bool {
	s.RLock()
	defer s.RUnlock()
	return s.data
}

// Add appends elem in Set s
func (s *set) Add(elem int) {
	s.Lock()
	defer s.Unlock()
	if s.data == nil {
		s.data = make(map[int]bool)
	}
	if s.data[elem] {
		return
	}
	s.data[elem] = true
}

// Remove deletes elem from Set s.
// Returns true, if delete is successfully, else return false
func (s *set) Remove(elem int) bool {
	s.Lock()
	defer s.Unlock()
	if !s.data[elem] {
		return false
	}
	delete(s.data, elem)
	return true
}

// Union unites two Sets this object and s2.
// Returns new Set
func (s *set) Union(s2 Set) Set {
	newS := NewSet()

	s.RLock()
	for item, _ := range s.data {
		newS.Add(item)
	}
	s.RUnlock()
	for item, _ := range s2.GetValues() {
		newS.Add(item)
	}

	return newS
}

// Difference returns new Set, contains elements of the this object, absent in the s2.
func (s *set) Difference(s2 Set) Set {
	newS := NewSet()

	s.RLock()
	defer s.RUnlock()
	for item, _ := range s.data {
		if !s2.GetValues()[item] {
			newS.Add(item)
		}
	}
	return newS
}

// Intersection returns new Set, containing common elements of a this object and a s2.
func (s *set) Intersection(s2 Set) Set {
	newS := NewSet()

	s.RLock()
	defer s.RUnlock()
	for item, _ := range s.data {
		if s2.GetValues()[item] {
			newS.Add(item)
		}
	}
	return newS
}

// Subset checks if a this object is a subset s2.
// Returns boolean.
func (s *set) Subset(s2 Set) bool {
	s.RLock()
	defer s.RUnlock()
	for item, _ := range s.data {
		if !s2.GetValues()[item] {
			return false
		}
	}
	return true
}

// NewSet creates and returns a new Set
func NewSet() Set {
	return &set{}
}
