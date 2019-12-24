package parentheses

// Validator is interface, contains method Validation
type Validator interface {
	Validation(string) bool
}

type parenthesesValidator struct {
	data map[string]string
}

// Parentheses checks the validity of a string of brackets.
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

// NewValidator creates new object, implements Validator, fills and returns his.
func NewValidator() Validator {
	obj := &parenthesesValidator{}
	obj.data = make(map[string]string)
	obj.data[")"] = "("
	obj.data["}"] = "{"
	obj.data["]"] = "["
	return obj
}
