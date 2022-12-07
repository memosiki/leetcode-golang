package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min0(a, b int) int {
	// 0-aware минимум
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	if a < b {
		return a
	}
	return b
}

/* Lists */

func SumList(container []int) (rolling_sum int) {
	for _, elem := range container {
		rolling_sum += elem
	}
	return
}

func MaxList(container []int) (rollingMax int) {
	for _, elem := range container {
		rollingMax = max(rollingMax, elem)
	}
	return
}

func AddLists(a, b []int) (sum []int) {
	// panics if b shorter than a
	sum = make([]int, len(a))
	for i := range sum {
		sum[i] = a[i] + b[i]
	}
	return
}

func ReverseList(container []int) {
	for i, j := 0, len(container)-1; i < j; i, j = i+1, j-1 {
		container[i], container[j] = container[j], container[i]
	}
}

func UniqueValues(container []int) (counter Set) {
	counter = make(Set)
	for _, value := range container {
		counter.Add(value)
	}
	return
}
