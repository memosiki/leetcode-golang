package main

import (
	"sort"
)

const MaxVal = 100_000

func longestSquareStreak(nums []int) (ans int) {
	seen := make([]bool, MaxVal+1)

	for _, num := range nums {
		seen[num] = true
	}
	sort.Ints(nums)
	var rolling int
	for _, num := range nums {
		rolling = 1
		for ; num*num <= MaxVal && seen[num*num]; num = num * num {
			seen[num] = false
			rolling++
		}
		ans = Max(ans, rolling)
	}
	if ans == 1 {
		return -1
	}
	return
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
