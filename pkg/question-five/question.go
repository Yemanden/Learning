package parentheses

// Validator is interface, contains method Validation
type Validator interface {
	Validation(string) bool
}

type parenthesesValidator struct {
	data map[string]string
}

// Validation checks the validity of a string of brackets.
// Returns a bool type
func (v *parenthesesValidator) Validation(input string) bool {
	stack := []string{}
	for i := 0; i < len(input); i++ {
		bracket := input[i : i+1]
		if v.data[bracket] == "" {
			stack = append(stack, bracket)
			continue
		}
		if len(stack) > 0 && stack[len(stack)-1] == v.data[bracket] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return true
}

// NewValidator creates a new object, implements Validator, fills and returns his.
func NewValidator() Validator {
	return &parenthesesValidator{map[string]string{")": "(", "}": "{", "]": "["}}
}
