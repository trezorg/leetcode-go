package sirsa

import "testing"

func TestFindMinInRotatedSortedArray(t *testing.T) {

	s := []int{1, 2, 3, -1, 0}
	expected := -1
	minEl := s[findMin(s)]
	if minEl != expected {
		t.Fatalf("Expected %d but got %d. Array: %#v", expected, minEl, s)
	}

	s = []int{1, 3}
	expected = 1
	minEl = s[findMin(s)]
	if minEl != expected {
		t.Fatalf("Expected %d but got %d. Array: %#v", expected, minEl, s)
	}

	s = []int{-1, 0, 1, 2, 3}
	expected = -1
	minEl = s[findMin(s)]
	if minEl != expected {
		t.Fatalf("Expected %d but got %d. Array: %#v", expected, minEl, s)
	}

	s = []int{-2, -1, 0, 1, 2, 3}
	expected = -2
	minEl = s[findMin(s)]
	if minEl != expected {
		t.Fatalf("Expected %d but got %d. Array: %#v", expected, minEl, s)
	}

	s = []int{-6, -1, 0, 1, 2, 3}
	expected = -6
	minEl = s[findMin(s)]
	if minEl != expected {
		t.Fatalf("Expected %d but got %d. Array: %#v", expected, minEl, s)
	}

	s = []int{0, 1, 2, 3, -6, -1}
	expected = -6
	minEl = s[findMin(s)]
	if minEl != expected {
		t.Fatalf("Expected %d but got %d. Array: %#v", expected, minEl, s)
	}

	s = []int{12, 15, 16, 17, 0, 1, 2, 3, 4, 5, 6}
	expected = 0
	minEl = s[findMin(s)]
	if minEl != expected {
		t.Fatalf("Expected %d but got %d. Array: %#v", expected, minEl, s)
	}

	s = []int{12, 15, 16, 0, 1, 2, 3}
	expected = 0
	minEl = s[findMin(s)]
	if minEl != expected {
		t.Fatalf("Expected %d but got %d. Array: %#v", expected, minEl, s)
	}

	s = []int{12, 15, 16, -1, 1, 2, 3, 4, 5, 6, 7}
	expected = -1
	minEl = s[findMin(s)]
	if minEl != expected {
		t.Fatalf("Expected %d but got %d. Array: %#v", expected, minEl, s)
	}

}

func TestFindMinInRotatedSortedArrayWithDuplicates(t *testing.T) {

	s := []int{1, 1, 1, 3, 1}
	expected := 4
	minIdx := findMinPivot(s)
	if minIdx != expected {
		t.Fatalf("Expected %d but got %d. Array: %#v", expected, minIdx, s)
	}

	s = []int{1, 3, 1, 1, 1}
	expected = 2
	minIdx = findMinPivot(s)
	if minIdx != expected {
		t.Fatalf("Expected %d but got %d. Array: %#v", expected, minIdx, s)
	}

	s = []int{1, 1, 3, 1, 1}
	expected = 3
	minIdx = findMinPivot(s)
	if minIdx != expected {
		t.Fatalf("Expected %d but got %d. Array: %#v", expected, minIdx, s)
	}

	s = []int{1, 1, 1, 1, 3}
	expected = 0
	minIdx = findMinPivot(s)
	if minIdx != expected {
		t.Fatalf("Expected %d but got %d. Array: %#v", expected, minIdx, s)
	}

}

func TestFindInRotatedSortedArray(t *testing.T) {

	s := []int{1, 2, 3, -1, 0}
	target := 3
	expected := 2
	result := search(s, target)
	if result != expected {
		t.Fatalf("Expected %d but got %d. Array: %#v. Target: %d", expected, result, s, target)
	}

	s = []int{}

	target = 3
	expected = -1
	result = search(s, target)
	if result != expected {
		t.Fatalf("Expected %d but got %d. Array: %#v. Target: %d", expected, result, s, target)
	}

	s = []int{1, 3}

	target = 0
	expected = -1
	result = search(s, target)
	if result != expected {
		t.Fatalf("Expected %d but got %d. Array: %#v. Target: %d", expected, result, s, target)
	}

	s = []int{1, 3}

	target = 1
	expected = 0
	result = search(s, target)
	if result != expected {
		t.Fatalf("Expected %d but got %d. Array: %#v. Target: %d", expected, result, s, target)
	}

	s = []int{1, 3}

	target = 4
	expected = -1
	result = search(s, target)
	if result != expected {
		t.Fatalf("Expected %d but got %d. Array: %#v. Target: %d", expected, result, s, target)
	}

	s = []int{1, 3, 5}

	target = 6
	expected = -1
	result = search(s, target)
	if result != expected {
		t.Fatalf("Expected %d but got %d. Array: %#v. Target: %d", expected, result, s, target)
	}

	s = []int{1, 3, 5}

	target = 0
	expected = -1
	result = search(s, target)
	if result != expected {
		t.Fatalf("Expected %d but got %d. Array: %#v. Target: %d", expected, result, s, target)
	}

	s = []int{1, 3, 5}

	target = 6
	expected = -1
	result = search(s, target)
	if result != expected {
		t.Fatalf("Expected %d but got %d. Array: %#v. Target: %d", expected, result, s, target)
	}

	s = []int{1}

	target = 3
	expected = -1
	result = search(s, target)
	if result != expected {
		t.Fatalf("Expected %d but got %d. Array: %#v. Target: %d", expected, result, s, target)
	}

	s = []int{12, 15, 16, -1, 1, 2, 3, 4, 5, 6, 7}

	target = 3
	expected = 6
	result = search(s, target)
	if result != expected {
		t.Fatalf("Expected %d but got %d. Array: %#v. Target: %d", expected, result, s, target)
	}

	target = 12
	expected = 0
	result = search(s, target)
	if result != expected {
		t.Fatalf("Expected %d but got %d. Array: %#v. Target: %d", expected, result, s, target)
	}

	target = 7
	expected = 10
	result = search(s, target)
	if result != expected {
		t.Fatalf("Expected %d but got %d. Array: %#v. Target: %d", expected, result, s, target)
	}

	target = -1
	expected = 3
	result = search(s, target)
	if result != expected {
		t.Fatalf("Expected %d but got %d. Array: %#v. Target: %d", expected, result, s, target)
	}

	target = 1
	expected = 4
	result = search(s, target)
	if result != expected {
		t.Fatalf("Expected %d but got %d. Array: %#v. Target: %d", expected, result, s, target)
	}

	target = 0
	expected = -1
	result = search(s, target)
	if result != expected {
		t.Fatalf("Expected %d but got %d. Array: %#v. Target: %d", expected, result, s, target)
	}

}

func TestFindInRotatedSortedArray2(t *testing.T) {

	s := []int{1, 2, 3, -1, 0}
	target := 3
	expected := true
	result := search2(s, target)
	if result != expected {
		t.Fatalf("Expected %v but got %v. Array: %#v. Target: %d", expected, result, s, target)
	}

	s = []int{}

	target = 3
	expected = false
	result = search2(s, target)
	if result != expected {
		t.Fatalf("Expected %v but got %v. Array: %#v. Target: %d", expected, result, s, target)
	}

	s = []int{1, 3}

	target = 0
	expected = false
	result = search2(s, target)
	if result != expected {
		t.Fatalf("Expected %v but got %v. Array: %#v. Target: %d", expected, result, s, target)
	}

	s = []int{1, 3}

	target = 1
	expected = true
	result = search2(s, target)
	if result != expected {
		t.Fatalf("Expected %v but got %v. Array: %#v. Target: %d", expected, result, s, target)
	}

	s = []int{1, 3}

	target = 4
	expected = false
	result = search2(s, target)
	if result != expected {
		t.Fatalf("Expected %v but got %v. Array: %#v. Target: %d", expected, result, s, target)
	}

	s = []int{1, 3, 5}

	target = 6
	expected = false
	result = search2(s, target)
	if result != expected {
		t.Fatalf("Expected %v but got %v. Array: %#v. Target: %d", expected, result, s, target)
	}

	s = []int{1, 3, 5}

	target = 0
	expected = false
	result = search2(s, target)
	if result != expected {
		t.Fatalf("Expected %v but got %v. Array: %#v. Target: %d", expected, result, s, target)
	}

	s = []int{1, 3, 5}

	target = 6
	expected = false
	result = search2(s, target)
	if result != expected {
		t.Fatalf("Expected %v but got %v. Array: %#v. Target: %d", expected, result, s, target)
	}

	s = []int{1}

	target = 3
	expected = false
	result = search2(s, target)
	if result != expected {
		t.Fatalf("Expected %v but got %v. Array: %#v. Target: %d", expected, result, s, target)
	}

	s = []int{12, 15, 16, -1, 1, 2, 3, 4, 5, 6, 7}

	target = 3
	expected = true
	result = search2(s, target)
	if result != expected {
		t.Fatalf("Expected %v but got %v. Array: %#v. Target: %d", expected, result, s, target)
	}

	target = 12
	expected = true
	result = search2(s, target)
	if result != expected {
		t.Fatalf("Expected %v but got %v. Array: %#v. Target: %d", expected, result, s, target)
	}

	target = 7
	expected = true
	result = search2(s, target)
	if result != expected {
		t.Fatalf("Expected %v but got %v. Array: %#v. Target: %d", expected, result, s, target)
	}

	target = -1
	expected = true
	result = search2(s, target)
	if result != expected {
		t.Fatalf("Expected %v but got %v. Array: %#v. Target: %d", expected, result, s, target)
	}

	target = 1
	expected = true
	result = search2(s, target)
	if result != expected {
		t.Fatalf("Expected %v but got %v. Array: %#v. Target: %d", expected, result, s, target)
	}

	target = 0
	expected = false
	result = search2(s, target)
	if result != expected {
		t.Fatalf("Expected %v but got %v. Array: %#v. Target: %d", expected, result, s, target)
	}

}

func TestFindInRotatedSortedArrayWithDuplicates2(t *testing.T) {

	s := []int{1, 2, 3, -1, 0, 0, 0}
	target := 3
	expected := true
	result := search2(s, target)
	if result != expected {
		t.Fatalf("Expected %v but got %v. Array: %#v. Target: %d", expected, result, s, target)
	}

	s = []int{}

	target = 3
	expected = false
	result = search2(s, target)
	if result != expected {
		t.Fatalf("Expected %v but got %v. Array: %#v. Target: %d", expected, result, s, target)
	}

	s = []int{1, 1, 1, 1, 3, 3, 3, 3}

	target = 0
	expected = false
	result = search2(s, target)
	if result != expected {
		t.Fatalf("Expected %v but got %v. Array: %#v. Target: %d", expected, result, s, target)
	}

	s = []int{1, 1, 1, 1, 3, 3, 3, 3}

	target = 1
	expected = true
	result = search2(s, target)
	if result != expected {
		t.Fatalf("Expected %v but got %v. Array: %#v. Target: %d", expected, result, s, target)
	}

	s = []int{1, 1, 1, 1, 3, 3, 3, 3}

	target = 4
	expected = false
	result = search2(s, target)
	if result != expected {
		t.Fatalf("Expected %v but got %v. Array: %#v. Target: %d", expected, result, s, target)
	}

	s = []int{1, 3, 5}

	target = 6
	expected = false
	result = search2(s, target)
	if result != expected {
		t.Fatalf("Expected %v but got %v. Array: %#v. Target: %d", expected, result, s, target)
	}

	s = []int{1, 3, 5}

	target = 0
	expected = false
	result = search2(s, target)
	if result != expected {
		t.Fatalf("Expected %v but got %v. Array: %#v. Target: %d", expected, result, s, target)
	}

	s = []int{1, 3, 5}

	target = 6
	expected = false
	result = search2(s, target)
	if result != expected {
		t.Fatalf("Expected %v but got %v. Array: %#v. Target: %d", expected, result, s, target)
	}

	s = []int{1}

	target = 3
	expected = false
	result = search2(s, target)
	if result != expected {
		t.Fatalf("Expected %v but got %v. Array: %#v. Target: %d", expected, result, s, target)
	}

	s = []int{1, 1, 1, 3, 1}

	target = 3
	expected = true
	result = search2(s, target)
	if result != expected {
		t.Fatalf("Expected %v but got %v. Array: %#v. Target: %d", expected, result, s, target)
	}

	s = []int{12, 12, 12, 12, 15, 16, -1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 3, 3, 3, 3, 3, 4, 5, 6, 7}

	target = 3
	expected = true
	result = search2(s, target)
	if result != expected {
		t.Fatalf("Expected %v but got %v. Array: %#v. Target: %d", expected, result, s, target)
	}

	target = 12
	expected = true
	result = search2(s, target)
	if result != expected {
		t.Fatalf("Expected %v but got %v. Array: %#v. Target: %d", expected, result, s, target)
	}

	target = 7
	expected = true
	result = search2(s, target)
	if result != expected {
		t.Fatalf("Expected %v but got %v. Array: %#v. Target: %d", expected, result, s, target)
	}

	target = -1
	expected = true
	result = search2(s, target)
	if result != expected {
		t.Fatalf("Expected %v but got %v. Array: %#v. Target: %d", expected, result, s, target)
	}

	target = 1
	expected = true
	result = search2(s, target)
	if result != expected {
		t.Fatalf("Expected %v but got %v. Array: %#v. Target: %d", expected, result, s, target)
	}

	target = 0
	expected = false
	result = search2(s, target)
	if result != expected {
		t.Fatalf("Expected %v but got %v. Array: %#v. Target: %d", expected, result, s, target)
	}

}
