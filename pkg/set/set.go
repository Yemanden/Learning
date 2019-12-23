package set

import "sync"

type Set interface {
	GetValues() map[int]bool
	Add(int)
	Remove(int) bool
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

// Union unites two Sets s1 and s2.
// Returns new Set
func Union(s1 Set, s2 Set) Set {
	numbers := make([]int, 0)

	for item, _ := range s1.GetValues() {
		numbers = append(numbers, item)
	}
	for item, _ := range s2.GetValues() {
		numbers = append(numbers, item)
	}

	return NewSet(numbers...)
}

// Difference returns new Set, contains elements of the s1, absent in the s2.
func Difference(s1 Set, s2 Set) Set {
	numbers := make([]int, 0)

	for item, _ := range s1.GetValues() {
		if !s2.GetValues()[item] {
			numbers = append(numbers, item)
		}
	}
	return NewSet(numbers...)
}

// Intersection returns new Set, containing common elements of a s1 and a s2.
func Intersection(s1 Set, s2 Set) Set {
	numbers := make([]int, 0)

	for item, _ := range s1.GetValues() {
		if s2.GetValues()[item] {
			numbers = append(numbers, item)
		}
	}
	return NewSet(numbers...)
}

// Subset checks if a s1 is a subset s2.
// Returns boolean.
func Subset(s1 Set, s2 Set) bool {
	for item, _ := range s1.GetValues() {
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
