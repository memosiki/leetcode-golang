package main

//import "fmt"

//func Max(a, b int) int {
//    if a > b {
//        return a
//    }
//    return b
//}

func MaxList(container []int) (rollingMax int) {
	for _, elem := range container {
		rollingMax = Max(rollingMax, elem)
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

type Set map[int]struct{}

func (s Set) Add(val int) {
	s[val] = struct{}{}
}

func (s Set) In(val int) bool {
	_, ok := s[val]
	return ok
}

func (s Set) Remove(val int) {
	delete(s, val)
}

func UniqueValues(container []int) (counter Set) {
	counter = make(Set)
	for _, value := range container {
		counter.Add(value)
	}
	return
}
func CountingSort(nums []int) {
	var counts = make([]int, MaxList(nums)+1)
	for _, elem := range nums {
		counts[elem] += 1
	}
	var i int
	for number, count := range counts {
		for k := 0; k < count; k++ {
			nums[i+k] = number
		}
		i += count
	}
}

func distinctAverages(nums []int) int {
	n := len(nums)
	CountingSort(nums)
	maxs, mins := nums[:n/2], nums[n/2:]
	ReverseList(maxs)
	return len(UniqueValues(AddLists(mins, maxs)))
}

//
//func main() {
//    fmt.Println(distinctAverages([]int{4,1,4,0,3,5}))
//}
