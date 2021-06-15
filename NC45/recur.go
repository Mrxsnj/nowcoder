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
	if root != nil {
		preRes = append(preRes, root.Val)
		pre(root.Left)
		pre(root.Right)
	}
}

func in(root *TreeNode) {
	if root != nil {
		in(root.Left)
		inRes = append(inRes, root.Val)
		in(root.Right)
	}
}

func post(root *TreeNode) {
	if root != nil {
		post(root.Left)
		post(root.Right)
		postRes = append(postRes, root.Val)
	}
}
