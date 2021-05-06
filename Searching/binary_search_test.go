package searching

import "testing"

func TestBinarySearch(t *testing.T) {

	arr := []int{1, 4, 5, 7, 8, 9, 12, 15, 29, 33}
	targets := []int{4, 1, 29, 27}

	want := []int{1, 0, 8, -1}

	for i, target := range targets {

		if got := binarySearch(arr, target); got != want[i] {
			t.Errorf("want %v and got %v", want[i], got)
		}
	}
}
