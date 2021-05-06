/*
	LEETCODE - Ques-2 (Add Two Numbers) (https://leetcode.com/problems/add-two-numbers/)
	You are given two non-empty linked lists representing two non-negative integers. The digits are
	stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and
	return the sum as a linked list.
	You may assume the two numbers do not contain any leading zero, except the number 0 itself.

	Example :

	num1:		2------>4------>3
	num2:		5------>6------>4
			----------------------------
	output:     7------>0------>8

	Input: l1 = [2,4,3], l2 = [5,6,4]
	Output: [7,0,8]
	Explanation: 342 + 465 = 807.
*/

package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 	Logic:
 		loop till both of the list are empty and while looping keep adding both the numbers and keep a not of the carry and
		 check after the loop whether the carry is 0 or not.
		 At every loop connect the current loop node to previous node
*/
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	res := &ListNode{}
	dummy := res

	carry := 0
	for l1 != nil || l2 != nil {

		x := 0
		y := 0
		if l1 != nil {
			x = l1.Val
		}
		if l2 != nil {
			y = l2.Val
		}

		sum := x + y + carry
		carry = sum / 10
		temp := &ListNode{sum % 10, nil}
		dummy.Next = temp
		dummy = dummy.Next

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry != 0 {
		lastNode := &ListNode{carry, nil}
		dummy.Next = lastNode
	}

	return res.Next
}
