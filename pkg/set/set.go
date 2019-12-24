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
	return s.data
}

// Add appends elem in Set s
func (s *set) Add(elem int) {
	if s.data[elem] {
		return
	}
	s.RLock()
	s.data[elem] = true
	s.RUnlock()
}

// Remove deletes elem from Set s.
// Returns true, if delete is successfully, else return false
func (s *set) Remove(elem int) bool {
	if !s.data[elem] {
		return false
	}
	s.RLock()
	s.data[elem] = false
	s.RUnlock()
	return true
}

// Union unites two Sets this object and s2.
// Returns new Set
func (s *set) Union(s2 Set) Set {
	numbers := make([]int, 0)

	for item, _ := range s.data {
		numbers = append(numbers, item)
	}
	for item, _ := range s2.GetValues() {
		numbers = append(numbers, item)
	}

	return NewSet(numbers...)
}

// Difference returns new Set, contains elements of the this object, absent in the s2.
func (s *set) Difference(s2 Set) Set {
	numbers := make([]int, 0)

	for item, _ := range s.data {
		if !s2.GetValues()[item] {
			numbers = append(numbers, item)
		}
	}
	return NewSet(numbers...)
}

// Intersection returns new Set, containing common elements of a this object and a s2.
func (s *set) Intersection(s2 Set) Set {
	numbers := make([]int, 0)

	for item, _ := range s.data {
		if s2.GetValues()[item] {
			numbers = append(numbers, item)
		}
	}
	return NewSet(numbers...)
}

// Subset checks if a this object is a subset s2.
// Returns boolean.
func (s *set) Subset(s2 Set) bool {
	for item, _ := range s.data {
		if !s2.GetValues()[item] {
			return false
		}
	}
	return true
}

// NewSet creates and returns Set with sets value of numbers
func NewSet(numbers ...int) Set {
	var newS set

	newS.data = make(map[int]bool)

	for _, item := range numbers {
		newS.Add(item)
	}

	return &newS
}
