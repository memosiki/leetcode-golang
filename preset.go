package main

//func main() {
//	a := 0
//	println(a)
//}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var Tree = &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}}

func DFS(node *TreeNode) {
	if node == nil {
		return
	}
	DFS(node.Left)
	DFS(node.Right)
}

type Element = int

func Max(a ...Element) (max Element) {
	max = a[0]
	for _, elem := range a {
		if max < elem {
			max = elem
		}
	}
	return
}
func Min(a ...Element) (min Element) {
	min = a[0]
	for _, elem := range a {
		if min > elem {
			min = elem
		}
	}
	return
}
func Abs(a Element) Element {
	if a < 0 {
		return -a
	}
	return a
}

func Sum(a ...Element) (rolling Element) {
	for _, elem := range a {
		rolling += elem
	}
	return
}
func Max2(a, b Element) Element {
	if a > b {
		return a
	}
	return b
}
func Min2(a, b Element) Element {
	if a < b {
		return a
	}
	return b
}

const MOD = 1_000_000_000 + 7
