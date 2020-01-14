package merge

// Merger ...
type Merger interface {
	Merge([]int, int, []int, int)
}

type merger struct{}

// Merge merges nums1 and nums2 in nums1 and sorts them
func (m *merger) Merge(nums1 []int, n int, nums2 []int, k int) {
	var i = 0
	var j = 0

	for i < n && j < k {
		if nums1[n-(i+1)] >= nums2[k-(j+1)] {
			nums1[n+k-(i+j+1)] = nums1[n-(i+1)]
			i++
			continue
		}
		nums1[n+k-(i+j+1)] = nums2[k-(j+1)]
		j++
	}
	for j < k {
		nums1[k-(j+1)] = nums2[k-(j+1)]
		j++
	}
}

// NewMerger ...
func NewMerger() Merger {
	return &merger{}
}
