package set

import (
	"reflect"
	"testing"
)

func TestGetValues(t *testing.T) {
	set := CreateSet(1, 2, 3, 4, 5, 6, 7, 8, 9)

	got := set.GetValues()
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("%s: got %v want %v", "GetValues", got, want)
	}
}

func TestCreateSet(t *testing.T) {

	got := CreateSet(1, 2, 3, 3, 2, 4).GetValues()
	want := []int{1, 2, 3, 4}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("%s: got %v want %v", "CreateSet", got, want)
	}
}

func TestAdd(t *testing.T) {
	set := CreateSet(1, 2, 3, 4, 5, 6, 7, 9)

	t.Run("missing", func(t *testing.T) {
		set.Add(8)
		got := set.GetValues()
		want := []int{1, 2, 3, 4, 5, 6, 7, 9, 8}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("%s: got %v want %v", "Add missing", got, want)
		}
	})

	set = CreateSet(1, 2, 3, 4, 5, 6, 7, 9)

	t.Run("existing", func(t *testing.T) {
		set.Add(1)
		got := set.GetValues()
		want := []int{1, 2, 3, 4, 5, 6, 7, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("%s: got %v want %v", "Add existing", got, want)
		}
	})
}

func TestRemove(t *testing.T) {
	set := CreateSet(1, 2, 3, 4, 5, 6, 7, 9)

	t.Run("missing", func(t *testing.T) {
		got := set.Remove(8)
		want := false

		if got != want {
			t.Errorf("%s: got %v want %v", "Remove missing", got, want)
		}
	})
	t.Run("existing", func(t *testing.T) {
		got := set.Remove(1)
		want := true

		if got != want {
			t.Errorf("%s: got %v want %v", "Remove existing", got, want)
		}
	})
}

func TestUnion(t *testing.T) {
	set1 := CreateSet(1, 2, 3, 4)
	set2 := CreateSet(3, 4, 5, 6)

	got := Union(set1, set2).GetValues()
	want := []int{1, 2, 3, 4, 5, 6}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("%s: got %v want %v", "Union", got, want)
	}
}

func TestDifference(t *testing.T) {
	set1 := CreateSet(1, 2, 3, 4)
	set2 := CreateSet(3, 4, 5, 6)

	got := Difference(set1, set2).GetValues()
	want := []int{1, 2}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("%s: got %v want %v", "Difference", got, want)
	}
}

func TestIntersection(t *testing.T) {
	set1 := CreateSet(1, 2, 3, 4)
	set2 := CreateSet(3, 4, 5, 6)

	got := Intersection(set1, set2).GetValues()
	want := []int{3, 4}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("%s: got %v want %v", "Intersection", got, want)
	}
}

func TestSubset(t *testing.T) {
	set1 := CreateSet(1, 2, 3, 4)
	set2 := CreateSet(3, 4, 5, 6)
	set3 := CreateSet(1, 2, 3, 4, 5)

	t.Run("missing", func(t *testing.T) {
		got := Subset(set1, set2)
		want := false

		if got != want {
			t.Errorf("%s: got %v want %v", "Subset missing", got, want)
		}
	})
	t.Run("existing", func(t *testing.T) {
		got := Subset(set1, set3)
		want := true

		if got != want {
			t.Errorf("%s: got %v want %v", "Subset existing", got, want)
		}
	})
}
