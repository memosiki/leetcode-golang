package main

import (
	"fmt"
	"github.com/emirpasic/gods/stacks/arraystack"
)

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

const MAX_INT = 3*10_000 + 1
const MOD = 1_000_000_000 + 7

func sumSubarrayMins(arr []int) int {
	var total int64
	fmt.Println("orig ", arr)

	n := len(arr)
	leftBound := make([]int, n)
	rightBound := make([]int, n)
	stack := arraystack.New()
	/*

	   {1 5 2 3}
	      ^
	*/
	stack.Push(0)
	for i := 1; i < n; i++ {
		top, _ := stack.Peek()
		for !stack.Empty() && arr[top.(int)] > arr[i] {
			rightBound[top.(int)] = i - 1
			stack.Pop()
			top, _ = stack.Peek()
		}
		stack.Push(i)

	}
	for !stack.Empty() {
		top, _ := stack.Pop()
		rightBound[top.(int)] = n - 1
	}

	stack.Push(n - 1)
	for i := n - 2; i >= 0; i-- {
		top, _ := stack.Peek()
		for !stack.Empty() && arr[top.(int)] >= arr[i] {
			leftBound[top.(int)] = i + 1
			stack.Pop()
			top, _ = stack.Peek()
		}
		stack.Push(i)

	}
	for !stack.Empty() {
		top, _ := stack.Pop()
		leftBound[top.(int)] = 0
	}

	fmt.Println("right", rightBound)
	fmt.Println("left ", leftBound)
	for i := range arr {
		total += int64((i - leftBound[i] + 1) * (rightBound[i] - i + 1) * arr[i])
		for total > MOD {
			total -= MOD
		}
	}

	return int(total)
}

//
//func main() {
//	a := sumSubarrayMins(
//		[]int{2, 2, 2, 2, 2},
//	)
//	fmt.Println(a)
//}
