package main

import . "nc_tools"

/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

/**
 *
 * @param pHead ListNodeš▒╗
 * @return ListNodeš▒╗
 */
func ReverseList(pHead *ListNode) *ListNode {
	// write code here
	var pre *ListNode
	var cur *ListNode
	var next *ListNode

	cur = pHead

	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre

}
