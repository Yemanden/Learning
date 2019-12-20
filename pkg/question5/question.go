package parentheses

// Parentheses checks the validity of a string of brackets.
// Returns a bool type
func Parentheses(input string) bool {
	stack := []string{}
	for i := 0; i < len(input); i++ {
		bracket := input[i : i+1]
		switch bracket {
		case ")":
			{
				if len(stack) > 0 && stack[len(stack)-1] == "(" {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			}
		case "}":
			{
				if len(stack) > 0 && stack[len(stack)-1] == "{" {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			}
		case "]":
			{
				if len(stack) > 0 && stack[len(stack)-1] == "[" {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			}
		default:
			{
				stack = append(stack, bracket)
				continue
			}
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}
