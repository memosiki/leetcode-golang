package main

import (
	"sort"
)

func main() {
	a := maxStarSum([]int{1, 2, 3, 4, 10, -10, -20}, [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}, {3, 5}, {3, 6}}, 2)
	println(a)
}

type tuple struct {
	from, val int
}

var starEdges [200_000]tuple

func maxStarSum(vals []int, edges [][]int, k int) (ans int) {
	var n int
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		if vals[to] > 0 {
			starEdges[n].val = vals[to]
			starEdges[n].from = from
			n++
		}
		if vals[from] > 0 {
			starEdges[n].val = vals[from]
			starEdges[n].from = to
			n++
		}
	}
	stars := starEdges[:n]
	sort.Slice(stars, func(i, j int) bool {
		return (stars[i].from == stars[j].from) && stars[i].val > stars[j].val || stars[i].from > stars[j].from
	})
	ans = Max(vals...)
	for i := 0; i < n; {
		star := stars[i]
		rolling := vals[star.from]
		for j := 0; j < k && i < n && stars[i].from == star.from; i, j = i+1, j+1 {
			rolling += stars[i].val
		}
		for ; i < n && stars[i].from == star.from; i++ {
		}
		ans = Max(ans, rolling)
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
