package merge_sorted_array

import (
	"sort"
)

// Merge merges nums1 and nums2 in nums1 and sorts them
func Merge(nums1 []int, n int, nums2 []int, m int) {
	for i := n; i < n+m; i++ {
		nums1[i] = nums2[i-n]
	}
	sort.Ints(nums1)
}
