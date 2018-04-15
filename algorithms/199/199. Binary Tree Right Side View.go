package algorithms

import "github.com/ljun20160606/leetcode/algorithms"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode = algorithms.TreeNode

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	var level []int
	queue := []*TreeNode{root}
	levelLen := len(queue)
	for len(queue) != 0 {
		if levelLen == 0 {
			result = append(result, level[len(level)-1])
			levelLen = len(queue)
			level = []int{}
		}
		t := queue[0]
		queue = queue[1:]
		level = append(level, t.Val)
		levelLen--
		if t.Left != nil {
			queue = append(queue, t.Left)
		}
		if t.Right != nil {
			queue = append(queue, t.Right)
		}
	}
	if len(level) > 0 {
		result = append(result, level[len(level)-1])
	}
	return result
}
