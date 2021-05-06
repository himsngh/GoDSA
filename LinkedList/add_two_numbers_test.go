package linkedlist

import "testing"

func TestAddTwoNumbers(t *testing.T) {

	list1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	list2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}

	want := &ListNode{7, &ListNode{0, &ListNode{8, nil}}}

	if got := AddTwoNumbers(list1, list2); got.Val != want.Val || got.Next.Val != want.Next.Val ||
		got.Next.Next.Val != want.Next.Next.Val {
		t.Errorf("wanted %v and got %+v", want, got)
	}
}
