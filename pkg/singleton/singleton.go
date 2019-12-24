package singleton

import "sync"

// Singleton is interface, contains methods Add and GetData
type Singleton interface {
	Add(int)
	GetData() int
}

// Singleton
type singleton struct {
	sync.RWMutex
	data int
}

// Global variable for singleton
var instance *singleton

// GetInstance returns Singleton object
func GetInstance() Singleton {
	if instance == nil {
		instance = &singleton{}
		instance.RLock()
		instance.data = 0
		instance.RUnlock()
	}
	return instance
}

// Add increases value of variable data in Singleton by i
func (s *singleton) Add(i int) {
	s.RLock()
	s.data += i
	s.RUnlock()
}

// GetData returns value of data in Singleton
func (s *singleton) GetData() int {
	s.RLock()
	ret := s.data
	s.RUnlock()
	return ret
}
