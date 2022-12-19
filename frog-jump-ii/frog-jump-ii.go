package main

func main() {
	a := maxJump([]int{0, 3, 9})
	println(a)
}
func maxJump(stones []int) (cost int) {
	n := len(stones)
	//	diffs := make([]int, n)
	var i int
	i = 2
	for ; i < n; i = i + 2 {
		cost = Max(cost, stones[i]-stones[i-2])
	}
	if i != n-1 {
		cost = Max(cost, stones[n-1]-stones[i-2])
	}
	i = 3
	cost = Max(cost, stones[0]-stones[1])
	for ; i < n; i = i + 2 {
		cost = Max(cost, stones[i]-stones[i-2])
	}
	if i != n-1 {
		cost = Max(cost, stones[n-1]-stones[i-2])
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
