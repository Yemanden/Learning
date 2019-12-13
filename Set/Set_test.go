package Set

import (
	"reflect"
	"testing"
)

func TestGetValues(t *testing.T) {
	set := CreateSet(1, 2, 3, 4, 5, 6, 7, 8, 9)

	got := set.GetValues()
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("%v: got %v want %v", t, got, want)
	}
}

func TestCreateSet(t *testing.T) {

	got := CreateSet(1, 2, 3, 3, 2, 4).GetValues()
	want := []int{1, 2, 3, 4}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("%v: got %v want %v", t, got, want)
	}
}

func TestAdd(t *testing.T) {
	set := CreateSet(1, 2, 3, 4, 5, 6, 7, 9)

	t.Run("missing", func(t *testing.T) {
		set.Add(8)
		got := set.GetValues()
		want := []int{1, 2, 3, 4, 5, 6, 7, 9, 8}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("%v: got %v want %v", t, got, want)
		}
	})
<<<<<<< HEAD

	set = CreateSet(1, 2, 3, 4, 5, 6, 7, 9)

=======
<<<<<<< HEAD
<<<<<<< HEAD
=======

	set = CreateSet(1, 2, 3, 4, 5, 6, 7, 9)

>>>>>>> 47c004a... Исправление названия файла-теста.
=======

	set = CreateSet(1, 2, 3, 4, 5, 6, 7, 9)

=======
<<<<<<< HEAD
=======

	set = CreateSet(1, 2, 3, 4, 5, 6, 7, 9)

>>>>>>> 47c004a... Исправление названия файла-теста.
>>>>>>> parent 87360e2558227203918361c46afa63e478ed8cf8
>>>>>>> a3daf9b... parent 87360e2558227203918361c46afa63e478ed8cf8
>>>>>>> Добавление задачи первого дня и реализации Множеств с тестами
	t.Run("existing", func(t *testing.T) {
		set.Add(1)
		got := set.GetValues()
		want := []int{1, 2, 3, 4, 5, 6, 7, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("%v: got %v want %v", t, got, want)
		}
	})
}

func TestRemove(t *testing.T) {
	set := CreateSet(1, 2, 3, 4, 5, 6, 7, 9)

	t.Run("missing", func(t *testing.T) {
		got := set.Remove(8)
		want := false

		if got != want {
			t.Errorf("%v: got %v want %v", t, got, want)
		}
	})
	t.Run("existing", func(t *testing.T) {
		got := set.Remove(1)
		want := true

		if got != want {
			t.Errorf("%v: got %v want %v", t, got, want)
		}
	})
}

func TestUnion(t *testing.T) {
	set1 := CreateSet(1, 2, 3, 4)
	set2 := CreateSet(3, 4, 5, 6)

<<<<<<< HEAD
	got := Union(set1, set2).GetValues()
=======
<<<<<<< HEAD
<<<<<<< HEAD
=======
	got := Union(set1, set2).GetValues()
=======
<<<<<<< HEAD
>>>>>>> a3daf9b... parent 87360e2558227203918361c46afa63e478ed8cf8
	got := Union(set1, set2)
=======
	got := Union(set1, set2).GetValues()
>>>>>>> 47c004a... Исправление названия файла-теста.
<<<<<<< HEAD
=======
>>>>>>> parent 87360e2558227203918361c46afa63e478ed8cf8
>>>>>>> a3daf9b... parent 87360e2558227203918361c46afa63e478ed8cf8
>>>>>>> Добавление задачи первого дня и реализации Множеств с тестами
	want := []int{1, 2, 3, 4, 5, 6}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("%v: got %v want %v", t, got, want)
	}
}

func TestDifference(t *testing.T) {
	set1 := CreateSet(1, 2, 3, 4)
	set2 := CreateSet(3, 4, 5, 6)

<<<<<<< HEAD
	got := Difference(set1, set2).GetValues()
=======
<<<<<<< HEAD
<<<<<<< HEAD
	got := Difference(set1, set2)
=======
	got := Difference(set1, set2).GetValues()
>>>>>>> 47c004a... Исправление названия файла-теста.
=======
	got := Difference(set1, set2).GetValues()
=======
<<<<<<< HEAD
	got := Difference(set1, set2)
=======
	got := Difference(set1, set2).GetValues()
>>>>>>> 47c004a... Исправление названия файла-теста.
>>>>>>> parent 87360e2558227203918361c46afa63e478ed8cf8
>>>>>>> a3daf9b... parent 87360e2558227203918361c46afa63e478ed8cf8
>>>>>>> Добавление задачи первого дня и реализации Множеств с тестами
	want := []int{1, 2}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("%v: got %v want %v", t, got, want)
	}
}

<<<<<<< HEAD
=======
<<<<<<< HEAD
<<<<<<< HEAD
=======
=======
<<<<<<< HEAD
>>>>>>> a3daf9b... parent 87360e2558227203918361c46afa63e478ed8cf8
func TestIntersection(t testing.T) {
	set1 := CreateSet(1, 2, 3, 4)
	set2 := CreateSet(3, 4, 5, 6)

	got := Intersection(set1, set2)
=======
<<<<<<< HEAD
=======
>>>>>>> parent 87360e2558227203918361c46afa63e478ed8cf8
>>>>>>> a3daf9b... parent 87360e2558227203918361c46afa63e478ed8cf8
>>>>>>> Добавление задачи первого дня и реализации Множеств с тестами
func TestIntersection(t *testing.T) {
	set1 := CreateSet(1, 2, 3, 4)
	set2 := CreateSet(3, 4, 5, 6)

	got := Intersection(set1, set2).GetValues()
<<<<<<< HEAD
=======
<<<<<<< HEAD
>>>>>>> 47c004a... Исправление названия файла-теста.
=======
<<<<<<< HEAD
=======
>>>>>>> 47c004a... Исправление названия файла-теста.
>>>>>>> parent 87360e2558227203918361c46afa63e478ed8cf8
>>>>>>> a3daf9b... parent 87360e2558227203918361c46afa63e478ed8cf8
>>>>>>> Добавление задачи первого дня и реализации Множеств с тестами
	want := []int{3, 4}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("%v: got %v want %v", t, got, want)
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
<<<<<<< HEAD
=======
<<<<<<< HEAD
<<<<<<< HEAD
=======
=======
<<<<<<< HEAD
>>>>>>> a3daf9b... parent 87360e2558227203918361c46afa63e478ed8cf8
			t.Errorf("%v: got %v want %v", t, got, want)
		}
	})
	t.Run("existing", func(t *testing.T) {
		got := Subset(set3, set1)
		want := true

		if got != want {
			t.Errorf("%v: got %v want %v", t, got, want)
=======
<<<<<<< HEAD
=======
>>>>>>> parent 87360e2558227203918361c46afa63e478ed8cf8
>>>>>>> a3daf9b... parent 87360e2558227203918361c46afa63e478ed8cf8
>>>>>>> Добавление задачи первого дня и реализации Множеств с тестами
			t.Errorf("%s: got %v want %v", "Subset missing", got, want)
		}
	})
	t.Run("existing", func(t *testing.T) {
		got := Subset(set1, set3)
		want := true

		if got != want {
			t.Errorf("%s: got %v want %v", "Subset existing", got, want)
<<<<<<< HEAD
=======
<<<<<<< HEAD
>>>>>>> 47c004a... Исправление названия файла-теста.
=======
<<<<<<< HEAD
=======
>>>>>>> 47c004a... Исправление названия файла-теста.
>>>>>>> parent 87360e2558227203918361c46afa63e478ed8cf8
>>>>>>> a3daf9b... parent 87360e2558227203918361c46afa63e478ed8cf8
>>>>>>> Добавление задачи первого дня и реализации Множеств с тестами
		}
	})
}
