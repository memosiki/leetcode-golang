package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	a := maxProduct(Tree)
	println(a)
}

var Tree = &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}}

type Element = int

func maxProduct(root *TreeNode) (ans int) {
	overall := DFSSum(root)
	var DFSSearch func(node *TreeNode) (split, size int)
	DFSSearch = func(node *TreeNode) (split, size int) {
		if node == nil {
			return overall, 0
		}
		splitL, sizeL := DFSSearch(node.Left)
		splitR, sizeR := DFSSearch(node.Right)

		size = sizeL + sizeR + node.Val
		split = size

		if Abs(overall-2*splitL) < Abs(overall-2*split) {
			split = splitL
		}
		if Abs(overall-2*splitR) < Abs(overall-2*split) {
			split = splitR
		}
		return
	}
	part1, _ := DFSSearch(root)
	ans = int(int64(part1) * int64(overall-part1) % MOD)
	return
}

func DFSSum(node *TreeNode) (sum int) {
	if node == nil {
		return
	}
	sum += node.Val
	sum += DFSSum(node.Left)
	sum += DFSSum(node.Right)
	return
}

func Abs(a Element) Element {
	if a < 0 {
		return -a
	}
	return a
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
func Max(a ...Element) (max Element) {
	max = a[0]
	for _, elem := range a {
		if max < elem {
			max = elem
		}
	}
	return
}
func Sum(a ...Element) (rolling Element) {
	for _, elem := range a {
		rolling += elem
	}
	return
}

const MOD = 1_000_000_000 + 7
