package questionthirteen

import "strconv"

// Calculator calculates...
type Calculator interface {
	Calculate([]int, []string) []int
}

// Calculate receive int's and strings arrays, returns int's array
func Calculate(numbers []int, signs []string) []int {
	result := make([]int, catalan(len(signs)))
	var iterator int
	for i := 0; i < len(signs); i++ {
		left := Calculate(numbers[:i+1], signs[:i])
		right := Calculate(numbers[i+1:], signs[i+1:])

		for j := 0; j < len(left); j++ {
			for k := 0; k < len(right); k++ {
				switch signs[i] {
				case "+":
					{
						result[iterator] = left[j] + right[k]
						iterator++
					}
				case "-":
					{
						result[iterator] = left[j] - right[k]
						iterator++
					}
				case "*":
					{
						result[iterator] = left[j] * right[k]
						iterator++
					}
				}
			}
		}
	}
	if len(signs) == 0 {
		result[0] = numbers[0]
	}
	return result
}

func diffWaysToCompute(input string) []int {
	var number string
	var signs []string
	var numbers []int
	for i := 0; i < len(input); i++ {
		if input[i:i+1] != "+" && input[i:i+1] != "-" && input[i:i+1] != "*" {
			number += input[i : i+1]
			continue
		}
		signs = append(signs, input[i:i+1])
		tmp, _ := strconv.Atoi(number)
		numbers = append(numbers, tmp)
		number = ""
	}
	tmp, _ := strconv.Atoi(number)
	numbers = append(numbers, tmp)
	return Calculate(numbers, signs)
}

func catalan(n int) int {
	var numerator = 1
	for i := n + 2; i <= (2*(n+2))-4; i++ {
		numerator *= i
	}
	return numerator / factorial(n)
}

func factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * factorial(i-1)
}
