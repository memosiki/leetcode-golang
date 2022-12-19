package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

const MaxNodes = 200
const Eof = -1

func dfs(channel chan<- int, node *TreeNode) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		channel <- node.Val
		return
	}
	dfs(channel, node.Left)
	dfs(channel, node.Right)
}
func startDfs(channel chan<- int, node *TreeNode) {
	dfs(channel, node)
	channel <- Eof
}
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	chan1 := make(chan int, MaxNodes)
	chan2 := make(chan int, MaxNodes)
	go startDfs(chan1, root1)
	go startDfs(chan2, root2)

	var leaf1, leaf2 int
	for {
		leaf1 = <-chan1
		leaf2 = <-chan2

		if leaf1 == Eof && leaf2 == Eof {
			return true
		} else if leaf1 != leaf2 {
			return false
		}
	}
}
