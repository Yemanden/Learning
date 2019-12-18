package Singleton

import "sync"

// Singleton
type singleton struct {
	m    sync.RWMutex
	data int
}

// Global variable for singleton
var instance *singleton

// Constructor of Singleton
func newSingleton() {
	instance = &singleton{}
	instance.m.RLock()
	instance.data = 0
	instance.m.RUnlock()
}

// GetInstance returns Singleton object
func GetInstance() *singleton {
	if instance == nil {
		newSingleton()
	}
	return instance
}

// Add increases value of variable data in Singleton by i
func (obj *singleton) Add(i int) {
	obj.m.RLock()
	obj.data += i
	obj.m.RUnlock()
}

// GetData returns value of data in Singleton
func (obj *singleton) GetData() int {
	obj.m.RLock()
	ret := obj.data
	obj.m.RUnlock()
	return ret
}
