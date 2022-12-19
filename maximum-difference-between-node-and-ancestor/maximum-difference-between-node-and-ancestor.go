package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const (
	MAX_INT = 100_001
	MIN_INT = -1
)

func Min(a ...int) (min int) {
	min = a[0]
	for _, elem := range a {
		if min > elem {
			min = elem
		}
	}
	return
}
func Max(a ...int) (max int) {
	max = a[0]
	for _, elem := range a {
		if max < elem {
			max = elem
		}
	}
	return
}
func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func DFS(node *TreeNode) (min int, max int, ans int) {
	if node == nil {
		return MAX_INT, MIN_INT, 0
	}
	minL, maxL, ansL := DFS(node.Left)
	minR, maxR, ansR := DFS(node.Right)
	min = Min(minL, minR, node.Val)
	max = Max(maxL, maxR, node.Val)
	ans = Max(ansL, ansR, Abs(node.Val-min), Abs(node.Val-max))
	return
}

func maxAncestorDiff(root *TreeNode) (ans int) {
	_, _, ans = DFS(root)
	return
}

func main() {
	a := maxAncestorDiff(&TreeNode{Right: &TreeNode{Val: 1}, Left: &TreeNode{Val: 3}})
	print(a)
}
