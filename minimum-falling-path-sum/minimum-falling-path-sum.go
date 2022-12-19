package main

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	prev := make([]int, n)
	cur := make([]int, n)
	copy(cur, matrix[0])

	var pathUp = func(i int) []int {
		start, end := i, i
		if 0 <= i-1 {
			start -= 1
		}
		if n > i+1 {
			end += 1
		}
		return prev[start : end+1]
	}
	for row := 1; row < n; row++ {
		prev, cur = cur, prev
		for i := 0; i < n; i++ {
			cur[i] = Min(pathUp(i)...) + matrix[row][i]
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
