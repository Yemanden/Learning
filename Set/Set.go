package Set

import "sync"

type set struct {
	mx   sync.RWMutex
	data []int
}

//Получить значения элементов Множества
func (s *set) GetValues() []int {
	return s.data
}

//Создать Множество с задаными значениями numbers (лишние значения отсеются)
func CreateSet(numbers ...int) *set {
	var newS set

	for _, item := range numbers {
		newS.Add(item)
	}

	return &newS
}

//Приватный метод поиска элемента elem в Множестве
func (s *set) find(elem int) (int, bool) {
	for i, item := range s.data {
		if item == elem {
			return i, true
		}
	}
	return 0, false
}

//Добавить элемент elem в Множество
func (s *set) Add(elem int) {
	_, ok := s.find(elem)
	if !ok {
		s.mx.RLock()
		s.data = append(s.data, elem)
		s.mx.RUnlock()
	}
}

//Удалить элемент elem из Множества
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

//Объединение Множеств s1 и s2.
//Возвращает новое Множество
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

//Разность Множеств s1 и s2.
//Возвращает Множество, состоящее из элементов Множества s1,
//которых нет в Множестве s2
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

//Пересечение Множеств s1 и s2.
//Возвращает Множество, состоящее из общих элементов
//Множеств s1 и s2
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

//Проверка на то, что Множество s1 является подмножеством
//множества s2.
//Возвращает bool
func Subset(s1 *set, s2 *set) bool {
	for _, item := range s1.data {
		_, ok := s2.find(item)
		if !ok {
			return false
		}
	}
	return true
}
