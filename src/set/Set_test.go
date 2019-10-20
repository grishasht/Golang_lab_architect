package set

import (
	"testing"
)

func TestItemSet_AddShouldAddCorrect(t *testing.T) {
	set := ItemSet{}

	set.Add(2)
	set.Add(4)
	set.Add(6)

	for _, x := range []int{2, 4, 6} {
		if !set.Contains(x) {
			t.Errorf("Set doesn't contains value %d", x)
		}
	}

	set.Add(4)
	set.Add(5)
	set.Add(6)

	for _, x := range []int{2, 4, 6, 5} {
		if !set.Contains(x) {
			t.Errorf("Set doesn't contains value %d", x)
		}
	}

	set.Add(7)

	result1 := set.Size()

	if result1 != 5 {
		t.Error("Expected 5, got ", result1)
	}
}

func TestItemSet_ClearShouldWorkCorrect(t *testing.T) {
	set := ItemSet{}
	set.Add(2)
	set.Add(4)
	set.Add(5)
	set.Add(6)

	set.Clear()

	result := set.Size()

	if result != 0 {
		t.Error("Expected 0, got ", result)
	}

	set.Add(1)
	set.Add(1)
	set.Add(1)

	set.Clear()

	result = set.Size()

	if result != 0 {
		t.Error("Expected 0, got ", result)
	}
}

func TestItemSet_DeleteShouldWorkCorrect(t *testing.T) {
	set := ItemSet{}
	set.Add(1)
	set.Add(2)
	set.Add(3)

	set.Delete(2)

	if set.Contains(2) {
		t.Errorf("Item %d must not exist in the set", 2)
	}
}

func TestItemSet_SizeShouldReturnSize3(t *testing.T) {
	set := ItemSet{}
	set.Add(1)
	set.Add(2)
	set.Add(3)

	result := set.Size()

	if result != 3 {
		t.Error("Expected 3, got ", result)
	}
}

func TestItemSet_SizeShouldReturnSize2(t *testing.T) {
	set := ItemSet{}
	set.Add(1)
	set.Add(2)
	set.Add(1)

	result := set.Size()

	if result != 2 {
		t.Error("Expected 2, got ", result)
	}
}

func TestItemSet_SizeShouldReturnCorrectSize(t *testing.T) {
	set := ItemSet{}
	set.Add(1)
	set.Add(2)
	set.Add(3)

	set.Delete(1)

	result := set.Size()

	if result != 2 {
		t.Error("Expected 2, got ", result)
	}
}

var set1 = ItemSet{}
var set2 = ItemSet{}

func TestItemSet_IntersectionShouldWorkCorrect(t *testing.T) {
	set1.Clear()
	set2.Clear()

	set1.Add(1)
	set1.Add(2)
	set1.Add(3)
	set1.Add(4)

	set2.Add(3)
	set2.Add(4)
	set2.Add(5)
	set2.Add(6)

	set3 := set1.Intersection(set2)

	if !set3.Has(3) || !set3.Has(4) {
		t.Error("Intersections doesn't works")
	}
}

func TestItemSet_DifferenceShouldWorkCorrect(t *testing.T) {
	set1.Clear()
	set2.Clear()

	set1.Add(1)
	set1.Add(2)
	set1.Add(3)
	set1.Add(4)

	set2.Add(3)
	set2.Add(4)
	set2.Add(5)
	set2.Add(6)

	set3 := set1.Difference(set2)

	if !set3.Has(1) || !set3.Has(2) {
		t.Error("Difference set1 & set2 doesn't contains elements: 1, 2")
	}

	set3 = set2.Difference(set1)

	if !set3.Has(5) || !set3.Has(6) {
		t.Error("Difference set2 & set1 doesn't contains elements: 5, 6")
	}
}

func TestItemSet_Union(t *testing.T) {
	set1.Clear()
	set2.Clear()

	set1.Add(1)
	set1.Add(2)
	set1.Add(3)
	set1.Add(4)

	set2.Add(3)
	set2.Add(4)
	set2.Add(5)
	set2.Add(6)

	set3 := set1.Union(set2)
	if !set3.Has(1) || !set3.Has(2) ||
		!set3.Has(4) || !set3.Has(3) ||
		!set3.Has(6) || !set3.Has(5) {
		t.Error("Union set1 & set2 doesn't contains all elements. Result: ", set3)
	}
}
