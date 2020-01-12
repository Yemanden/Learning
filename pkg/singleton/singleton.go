package singleton

import "sync"

type singletoner interface {
	Adder
	Getter
}

// Adder ...
type Adder interface {
	Add(int)
}

// Getter ...
type Getter interface {
	GetData() int
}

// singleton struct
type singleton struct {
	sync.RWMutex
	data int
}

// Global variable for singleton
var instance *singleton

// Add increases value of data in singleton by i
func (s *singleton) Add(i int) {
	s.Lock()
	s.data += i
	s.Unlock()
}

// GetData returns value of data in singleton
func (s *singleton) GetData() int {
	s.RLock()
	defer s.RUnlock()
	return s.data
}

// GetInstance returns ours singleton or initialized his
func GetInstance() singletoner {
	if instance == nil {
		instance = &singleton{}
		instance.Lock()
		defer instance.Unlock()
		instance.data = 0
	}
	return instance
}
