package main

import . "nc_tools"

/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

/**
 *
 * @param root TreeNode类 the root of binary tree
 * @return int整型二维数组
 */

var preRes []int
var inRes []int
var postRes []int

func threeOrders(root *TreeNode) [][]int {
	// write code here
	res := [][]int{}
	preRes = []int{}
	inRes = []int{}
	postRes = []int{}

	pre(root)
	res = append(res, preRes)

	in(root)
	res = append(res, inRes)

	post(root)
	res = append(res, postRes)

	return res
}

func pre(root *TreeNode) {
	st := []*TreeNode{}

	for {
		for root != nil {
			preRes = append(preRes, root.Val)
			st = append(st, root)
			root = root.Left
		}

		if len(st) == 0 {
			break
		} else {
			root = st[len(st)-1].Right
			st = st[:len(st)-1]
		}
	}
}

func in(root *TreeNode) {
	st := []*TreeNode{}

	for {
		for root != nil {
			st = append(st, root)
			root = root.Left
		}

		if len(st) == 0 {
			break
		} else {
			root = st[len(st)-1].Right
			inRes = append(inRes, st[len(st)-1].Val)
			st = st[:len(st)-1]
		}
	}
}

func post(root *TreeNode) {
	st := []*TreeNode{}
	var lastPrinted *TreeNode

	for {
		for root != nil {
			st = append(st, root)
			root = root.Left
		}

		if len(st) == 0 {
			break
		} else {
			root = st[len(st)-1]
		}

		if root.Right == nil || root.Right == lastPrinted {
			lastPrinted = root
			postRes = append(postRes, root.Val)
			st = st[:len(st)-1]
			root = nil
		} else {
			root = root.Right
		}
	}
}
