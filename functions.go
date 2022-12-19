package main

// Min0 0-aware минимум
func Min0(a, b Element) Element {
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

func Reverse(container []Element) {
	for i, j := 0, len(container)-1; i < j; i, j = i+1, j-1 {
		container[i], container[j] = container[j], container[i]
	}
}

// AddSlices return slice that is element-wise sum of two slices.
// Panics if b shorter than a
func AddSlices(a, b []Element) (sum []Element) {
	sum = make([]Element, len(a))
	for i := range sum {
		sum[i] = a[i] + b[i]
	}
	return
}

func UniqueValues(container []Element) (counter Set) {
	counter = make(Set)
	for _, value := range container {
		counter.Add(value)
	}
	return
}
