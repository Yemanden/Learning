package questionfour

import (
	"testing"
)

const (
	benchmarkTestMergePass1Name  = "10000 elements"
	benchmarkTestMergePass1Count = 10000

	benchmarkTestMergePass2Name  = "1000000 elements"
	benchmarkTestMergePass2Count = 1000000
)

// BenchmarkMerge ...
func BenchmarkMerge(b *testing.B) {
	b.Run(benchmarkTestMergePass1Name, func(b *testing.B) {
		b.StopTimer()

		nums1 := make([]int, 2*benchmarkTestMergePass1Count, 2*benchmarkTestMergePass1Count)
		for i := 0; i < benchmarkTestMergePass1Count; i++ {
			nums1[i] = i*2 + 3
		}

		nums2 := make([]int, benchmarkTestMergePass1Count)
		for i := 0; i < benchmarkTestMergePass1Count; i++ {
			nums2[i] = i*5 + 9
		}

		m := NewMerger()

		b.StartTimer()
		m.Merge(nums1, benchmarkTestMergePass1Count, nums2, benchmarkTestMergePass1Count)
	})

	b.Run(benchmarkTestMergePass2Name, func(b *testing.B) {
		b.StopTimer()

		nums1 := make([]int, 2*benchmarkTestMergePass2Count, 2*benchmarkTestMergePass2Count)
		for i := 0; i < benchmarkTestMergePass2Count; i++ {
			nums1[i] = i*2 + 3
		}

		nums2 := make([]int, benchmarkTestMergePass2Count)
		for i := 0; i < benchmarkTestMergePass2Count; i++ {
			nums2[i] = i*5 + 9
		}

		m := NewMerger()

		b.StartTimer()
		m.Merge(nums1, benchmarkTestMergePass2Count, nums2, benchmarkTestMergePass2Count)
	})
}
