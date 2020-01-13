package set

import "sync"

// Seter ...
type Seter interface {
	Getter
	Adder
	Remover
	Unioner
	Differencer
	Intersectioner
	IsSubset(Seter) bool
}

// Getter ...
type Getter interface {
	GetValues() map[int]bool
}

// Adder ...
type Adder interface {
	Add(int)
}

// Remover ...
type Remover interface {
	Remove(int) bool
}

// Unioner ...
type Unioner interface {
	Union(Seter) Seter
}

// Differencer ...
type Differencer interface {
	Difference(Seter) Seter
}

// Intersectioner ...
type Intersectioner interface {
	Intersection(Seter) Seter
}

type set struct {
	sync.RWMutex
	data map[int]bool
}

// GetValues returns values of the data contained in this Seter
func (s *set) GetValues() map[int]bool {
	s.RLock()
	defer s.RUnlock()
	return s.data
}

// Add appends elem in Seter s
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

// Remove deletes elem from Seter s.
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

// Union unites two Seters this object and s2.
// Returns new Seter
func (s *set) Union(s2 Seter) Seter {
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

// Difference returns new Seter, contains elements of the this object, absent in the s2.
func (s *set) Difference(s2 Seter) Seter {
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

// Intersection returns new Seter, containing common elements of a this object and a s2.
func (s *set) Intersection(s2 Seter) Seter {
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
func (s *set) IsSubset(s2 Seter) bool {
	s.RLock()
	defer s.RUnlock()
	for item, _ := range s.data {
		if s2.GetValues()[item] == false {
			return false
		}
	}
	return true
}

// NewSet creates and returns a new Seter
func NewSet() Seter {
	return &set{}
}
