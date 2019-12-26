package questionten

// Intersecter is interface, contains method Intersect
type Intersecter interface {
	Intersect([]int, []int) []int
}

type intersecter struct{}

// Intersect returns result of intersection two slices nums1 and nums2
func (i *intersecter) Intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for _, item := range nums1 {
		m[item]++
	}

	var answer []int
	for _, item := range nums2 {
		if m[item] > 0 {
			answer = append(answer, item)
			m[item]--
		}
	}

	return answer
}

// NewIntersecter returns a new object intersecter structure.
func NewIntersecter() Intersecter {
	return &intersecter{}
}
