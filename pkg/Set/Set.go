package Set

import "sync"

type set struct {
	mx   sync.RWMutex
	data []int
}

// GetValues возвращает значения элементов Множества
func (s *set) GetValues() []int {
	return s.data
}

// CreateSet создаёт и возвращает Множество с задаными значениями numbers (лишние значения отсеются)
func CreateSet(numbers ...int) *set {
	var newS set

	for _, item := range numbers {
		newS.Add(item)
	}

	return &newS
}

// Поиск элемента в множестве, возвращает индекс элемента и true, если найдено,
// иначе возвращает 0 и false.
func (s *set) find(elem int) (int, bool) {
	for i, item := range s.data {
		if item == elem {
			return i, true
		}
	}
	return 0, false
}

// Add добавляет элемент во Множество
func (s *set) Add(elem int) {
	_, ok := s.find(elem)
	if !ok {
		s.mx.RLock()
		s.data = append(s.data, elem)
		s.mx.RUnlock()
	}
}

// Remove удаляет заданный элемент из Множества и возвращает true, иначе false
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

// Union объединяет два Множества и возвращает результат объединения
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

// Difference возвращает Множество, состоящее из элементов
// множества s1, которых нет в множестве s2
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

// Intersection возвращает новое Множество, состоящее из общих элементов
// Множеств s1 и s2
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

// Subset проверяет является ли множество s1 подмножеством s2.
// Возвращает true, если является и false, если нет.
func Subset(s1 *set, s2 *set) bool {
	for _, item := range s1.data {
		_, ok := s2.find(item)
		if !ok {
			return false
		}
	}
	return true
}
