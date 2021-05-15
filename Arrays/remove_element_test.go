package arrays

import "testing"

func TestRemoveElement(t *testing.T) {

	array := []int{3, 2, 2, 3}
	val := 3

	want := 2

	if got := removeElement(array, val); got != want {
		t.Errorf("want length : %v,  got length : %v", want, got)
	}
}
