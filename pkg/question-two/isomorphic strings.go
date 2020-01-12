package questiontwo

// Isomorpher is interface, contains method Isomorphic
type Isomorpher interface {
	IsIsomorphic(s1, s2 string) bool
}

type isomorpher struct{}

// Isomorphic receives in parameters two strings, сhecks them for isomorphism
// and returns bool value
func (i *isomorpher) IsIsomorphic(s1, s2 string) bool {
	m1 := make(map[string]string, len(s1))
	m2 := make(map[string]string, len(s2))

	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if char, ok := m1[s1[i:i+1]]; ok && char != s2[i:i+1] {
			return false
		}
		m1[s1[i:i+1]] = s2[i : i+1]

		if char, ok := m2[s2[i:i+1]]; ok && char != s1[i:i+1] {
			return false
		}
		m2[s2[i:i+1]] = s1[i : i+1]
	}
	return true
}

// NewIsomorpher return a new object, implemented Isomorpher
func NewIsomorpher() Isomorpher {
	return &isomorpher{}
}
