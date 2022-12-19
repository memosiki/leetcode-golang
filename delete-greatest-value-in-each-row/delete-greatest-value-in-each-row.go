package main

import "sort"

//func main() {
//	a := 0
//	println(a)
//}

func deleteGreatestValue(grid [][]int) (ans int) {
	n := len(grid)
	m := len(grid[0])
	for i := range grid {
		sort.Ints(grid[i])
	}
	for j := 0; j < m; j++ {
		max := 0
		for i := 0; i < n; i++ {
			max = Max(max, grid[i][j])
		}
		ans += max
	}
	return
}

type Element = int

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
