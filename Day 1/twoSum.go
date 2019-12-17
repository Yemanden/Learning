package Day_1

// Метод для нахождения индексов двух элементов массива nums,
// сумма которых равна target
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ { // Получаем индексы каждого элемента
		for j := i + 1; j < len(nums); j++ { // Получаем индексы каждого следующего элемента
			if nums[i]+nums[j] == target { // Сравниваем сумму элементов с запрашиваемым значением
				return []int{i, j}
			}
		}
	}
	return nil
}
