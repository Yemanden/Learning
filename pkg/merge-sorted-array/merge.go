package merge_sorted_array

// Merge merges nums1 and nums2 in nums1 and sorts them
func merge(nums1 []int, n int, nums2 []int, m int) {
	var i int = 0
	var j int = 0

	for i < n && j < m {
		if nums1[n-(i+1)] >= nums2[m-(j+1)] {
			nums1[n+m-(i+j+1)] = nums1[n-(i+1)]
			i++
			continue
		}
		nums1[n+m-(i+j+1)] = nums2[m-(j+1)]
		j++
	}
	for j < m {
		nums1[m-(j+1)] = nums2[m-(j+1)]
		j++
	}
}
