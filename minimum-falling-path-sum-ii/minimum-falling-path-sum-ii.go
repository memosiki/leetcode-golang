package main

import "sort"

const MaxVal = 99
const MinVal = -99

func CountingSearchTop2(container []int) (min, top2 int) {
	foundMin := false
	counts := make([]int, -MinVal+MaxVal+1)
	for _, elem := range container {
		counts[-MinVal+elem]++
	}
	for num := MinVal; num < MaxVal; num++ {
		for k := 0; k < counts[-MinVal+num]; k++ {
			if foundMin {
				return min, num
			}
			foundMin = true
			min = num
		}
	}
	panic("Slice too short")
}

func minFallingPathSum(grid [][]int) int {
	n := len(grid)
	prev := make([]int, n)
	cur := make([]int, n)
	copy(cur, grid[0])

	var min, top2 int
	for row := 1; row < n; row++ {
		prev, cur = cur, prev
		//min, top2 = CountingSearchTop2(prev)
		copy(cur, prev)
		sort.Ints(cur)
		min, top2 = cur[0], cur[1]
		for i := 0; i < n; i++ {
			cur[i] = min + grid[row][i]
			if prev[i] == min {
				cur[i] = top2 + grid[row][i]
			}
		}
	}
	return Min(cur...)
}

type Element = int

func Min(a ...Element) (min Element) {
	min = a[0]
	for _, elem := range a {
		if min > elem {
			min = elem
		}
	}
	return
}
