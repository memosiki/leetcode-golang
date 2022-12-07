package main

import (
	"fmt"
	"github.com/emirpasic/gods/stacks/arraystack"
)

const MOD = 1_000_000_000 + 7

func totalStrength(strength []int) int {
	fmt.Println("orig ", strength)
	var total int64
	n := len(strength)
	leftBound := make([]int, n)
	rightBound := make([]int, n)
	prefixSum := make([]int, n)
	for i := 1; i < n; i++ {
		prefixSum[i] = prefixSum[i-1] + strength[i-1] // TODO: mod MOD
	}
	fmt.Println("pref ", prefixSum)

	stack := arraystack.New()
	/*

	   {1 5 2 3}
	   ^
	*/
	stack.Push(0)
	for i := 1; i < n; i++ {
		top, _ := stack.Peek()
		for !stack.Empty() && strength[top.(int)] > strength[i] {
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
		for !stack.Empty() && strength[top.(int)] >= strength[i] {
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
	for i := range strength {
		left := i - leftBound[i]
		right := rightBound[i] - i
		sum := prefixSum[rightBound[i]] - prefixSum[leftBound[i]] + strength[rightBound[i]]
		//		repeats := (left + 1) * (right + 1)
		total += int64(repeats * strength[i] * sum)
		fmt.Println(i, "->", strength[i], "(", left, ",", right, ")", "sum", sum, "times", repeats)
		for total > MOD {
			total -= MOD
		}
	}

	return int(total)
}

func main() {
	a := totalStrength(
		[]int{5, 4, 6},
	)
	fmt.Println(a)
}
