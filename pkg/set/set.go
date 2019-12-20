package set

import "sync"

type set struct {
	mx   sync.RWMutex
	data []int
}

// GetValues returns values of the data contained in this Set
func (s *set) GetValues() []int {
	return s.data
}

// CreateSet creates and returns Set with sets value of numbers
func CreateSet(numbers ...int) *set {
	var newS set

	for _, item := range numbers {
		newS.Add(item)
	}

	return &newS
}

func (s *set) find(elem int) (int, bool) {
	for i, item := range s.data {
		if item == elem {
			return i, true
		}
	}
	return 0, false
}

// Add appends elem in Set s
func (s *set) Add(elem int) {
	_, ok := s.find(elem)
	if !ok {
		s.mx.RLock()
		s.data = append(s.data, elem)
		s.mx.RUnlock()
	}
}

// Remove deletes elem from Set s.
// Returns true, if delete is successfully, else return false
func (s *set) Remove(elem int) bool {
	i, ok := s.find(elem)
	if ok {
		s.mx.RLock()
		s.data = append(s.data[:i], s.data[i+1:]...)
		s.mx.RUnlock()
		return true
	}
	return false
}

// Union unites two Sets s1 and s2.
// Returns new Set
func Union(s1 *set, s2 *set) *set {
	var newS set

	for _, item := range s1.data {
		newS.Add(item)
	}
	for _, item := range s2.data {
		newS.Add(item)
	}

	return &newS
}

// Difference returns new Set, contains elements of the s1, absent in the s2.
func Difference(s1 *set, s2 *set) *set {
	var newS set

	for _, item := range s1.data {
		_, ok := s2.find(item)
		if !ok {
			newS.Add(item)
		}
	}

	return &newS
}

// Intersection returns new Set, containing common elements of a s1 and a s2.
func Intersection(s1 *set, s2 *set) *set {
	var newS set

	for _, item := range s1.data {
		_, ok := s2.find(item)
		if ok {
			newS.Add(item)
		}
	}
	return &newS
}

// Subset checks if a s1 is a subset s2.
// Returns boolean.
func Subset(s1 *set, s2 *set) bool {
	for _, item := range s1.data {
		_, ok := s2.find(item)
		if !ok {
			return false
		}
	}
	return true
}
