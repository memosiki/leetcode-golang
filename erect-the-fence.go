package main

import (
	"fmt"
	"github.com/emirpasic/gods/stacks/arraystack"
	"sort"
)

func dot(a, b []int) int {
	return a[0]*b[0] + a[1]*b[1]
}
func outerTrees(trees [][]int) (answer [][]int) {
	var n = len(trees)
	sort.Slice(
		trees,
		func(i, j int) bool {
			return trees[i][0] == trees[j][0] && trees[i][1] < trees[j][1] || trees[i][0] < trees[j][0]
		},
	)
	fmt.Println(trees)
	stackUpper := arraystack.New()
	_ = stackUpper
	for _, treeB := range trees {
		for stackUpper.Size() > 0 {
			treeA, _ := stackUpper.Pop()
			if dot(treeA.([]int), treeB) >= 0 {
				break
			}
		}
		stackUpper.Push(treeB)
	}
	stackLower := arraystack.New()
	for i := n - 1; i > 0; i-- {
		treeB := trees[i]
		for stackLower.Size() > 0 {
			treeA, _ := stackLower.Pop()
			if dot(treeA.([]int), treeB) >= 0 {
				break
			}
		}
		stackLower.Push(treeB)
	}

	stackUpper.Pop()
	stackLower.Pop()
	for _, tree := range stackUpper.Values() {
		answer = append(answer, tree.([]int))
	}
	for _, tree := range stackLower.Values() {
		answer = append(answer, tree.([]int))
	}
	return
}

//func main() {
//	a := outerTrees(
//		[][]int{{1, 1}, {2, 2}, {2, 0}, {2, 4}, {3, 3}, {4, 2}},
//	)
//	fmt.Println(a)
//}
