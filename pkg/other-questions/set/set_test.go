package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testGetValueName = "GetValues"

	testAddPass1Name = "Add Missing"
	testAddPass2Name = "Add Existing"

	testRemovePass1Name = "Remove Missing"
	testRemovePass1Want = false

	testRemovePass2Name = "Remove Existing"
	testRemovePass2Want = true

	testUnionName = "Union"

	testDifferenceName = "Difference"

	testIntersectionName = "Intersection"

	testSubsetPass1Name = "Subset Missing"
	testSubsetPass1Want = false

	testSubsetPass2Name = "Subset Existing"
	testSubsetPass2Want = true
)

// TestGetValues ...
func TestGetValues(t *testing.T) {
	s := NewSet()
	for i := 1; i < 10; i++ {
		s.Add(i)
	}

	t.Run(testGetValueName, func(t *testing.T) {
		got := s.GetValues()
		want := map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true}

		assert.EqualValues(t, got, want)
	})
}

// TestAdd ...
func TestAdd(t *testing.T) {
	s := NewSet()
	for i := 1; i < 10; i++ {
		s.Add(i)
	}

	t.Run(testAddPass1Name, func(t *testing.T) {
		s.Add(10)
		got := s.GetValues()
		want := map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true, 10: true}

		assert.EqualValues(t, got, want)
	})

	s = NewSet()
	for i := 1; i < 10; i++ {
		s.Add(i)
	}

	t.Run(testAddPass2Name, func(t *testing.T) {
		s.Add(1)
		got := s.GetValues()
		want := map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true}

		assert.EqualValues(t, got, want)
	})
}

// TestRemove ...
func TestRemove(t *testing.T) {
	s := NewSet()
	for i := 1; i < 10; i++ {
		s.Add(i)
	}

	t.Run(testRemovePass1Name, func(t *testing.T) {
		got := s.Remove(10)
		want := testRemovePass1Want

		assert.EqualValues(t, got, want)
	})
	t.Run(testRemovePass2Name, func(t *testing.T) {
		got := s.Remove(1)
		want := testRemovePass2Want

		assert.EqualValues(t, got, want)
	})
}

// TestUnion ...
func TestUnion(t *testing.T) {
	s1 := NewSet()
	s2 := NewSet()
	for i := 1; i < 4; i++ {
		s1.Add(i)
		s2.Add(2 * i)
	}

	t.Run(testUnionName, func(t *testing.T) {
		got := s1.Union(s2).GetValues()
		want := map[int]bool{1: true, 2: true, 3: true, 4: true, 6: true}

		assert.EqualValues(t, got, want)
	})
}

// TestDifference ...
func TestDifference(t *testing.T) {
	s1 := NewSet()
	s2 := NewSet()
	for i := 1; i < 4; i++ {
		s1.Add(i)
		s2.Add(2 * i)
	}

	t.Run(testDifferenceName, func(t *testing.T) {
		got := s1.Difference(s2).GetValues()
		want := map[int]bool{1: true, 3: true}

		assert.EqualValues(t, got, want)
	})
}

// TestIntersection ...
func TestIntersection(t *testing.T) {
	s1 := NewSet()
	s2 := NewSet()
	for i := 1; i < 4; i++ {
		s1.Add(i)
		s2.Add(2 * i)
	}

	t.Run(testIntersectionName, func(t *testing.T) {
		got := s1.Intersection(s2).GetValues()
		want := map[int]bool{2: true}

		assert.EqualValues(t, got, want)
	})
}

// TestSubset ...
func TestSubset(t *testing.T) {
	s1 := NewSet()
	s2 := NewSet()
	for i := 1; i < 4; i++ {
		s1.Add(i)
		s2.Add(2 * i)
	}
	s3 := NewSet()
	for i := 1; i < 10; i++ {
		s3.Add(i)
	}

	t.Run(testSubsetPass1Name, func(t *testing.T) {
		got := s1.IsSubset(s2)
		want := testSubsetPass1Want

		assert.EqualValues(t, got, want)
	})
	t.Run(testSubsetPass2Name, func(t *testing.T) {
		got := s1.IsSubset(s3)
		want := testSubsetPass2Want

		assert.EqualValues(t, got, want)
	})
}
