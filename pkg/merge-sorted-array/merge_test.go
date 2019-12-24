package merge_sorted_array

import (
	"testing"
)

func BenchmarkMerge(b *testing.B) {

	for k := 0; k < 2; k++ {
		b.StopTimer()
		Count := 10000 + k*990000 // Два теста с 10000 и с 1000000 элементами

		nums1 := make([]int, 2*Count, 2*Count)
		for i := 0; i < Count; i++ {
			nums1[i] = i*2 + 3
		}

		nums2 := make([]int, Count)
		for i := 0; i < Count; i++ {
			nums2[i] = i*5 + 9
		}

		b.StartTimer()
		merge(nums1, Count, nums2, Count)
	}
}
