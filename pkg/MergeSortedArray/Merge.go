package Merge_Sorted_Array

import (
	"sort"
)

// Merge производит сортировочное слияние двух отсортированных массивов
// и записывает результат в первый массив
func Merge(nums1 []int, n int, nums2 []int, m int) {
	for i := n; i < n+m; i++ {
		nums1[i] = nums2[i-n]
	}
	sort.Ints(nums1)
}
