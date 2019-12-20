package question_two

//Функция, проверяющая изоморфность двух строк
func Isomorphic(s1 string, s2 string) bool {
	m1 := make(map[string]string)
	m2 := make(map[string]string)

	//Если длины строк разные, то строки заведомо неизоморфны
	if len(s1) != len(s2) {
		return false
	}

	//Цикл по всем элементам строк
	for i := 0; i < len(s1); i++ {
		//Часть, проверяющая, чтобы каждой уникальной букве первой строки соответствовала буква
		//из второй строки
		char, ok := m1[s1[i:i+1]]
		if ok && char != s2[i:i+1] {
			return false
		} else {
			m1[s1[i:i+1]] = s2[i : i+1]
		}

		//Часть, проверяющая, чтобы каждой уникальной букве вторй строки соответствовала буква
		//из первой строки
		char, ok = m2[s2[i:i+1]]
		if ok && char != s1[i:i+1] {
			return false
		} else {
			m2[s2[i:i+1]] = s1[i : i+1]
		}
	}
	return true
}
