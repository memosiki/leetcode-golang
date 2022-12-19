package main

import (
	"github.com/emirpasic/gods/sets/hashset"
	"math/bits"
	"sort"
)

var memory = [100_000]int{}

func countExcellentPairs(nums []int, k int) (answer int64) {
	var seen = hashset.New()
	var n int
	for _, num := range nums {
		if !seen.Contains(num) {
			memory[n] = bits.OnesCount(uint(num))
			seen.Add(num)
			n++
		}
	}
	bitVals := memory[:n]

	sort.Ints(bitVals)

	for _, bitVal := range bitVals {
		answer += int64(n - sort.SearchInts(bitVals, k-bitVal))
	}
	return
}

//func main() {
//	a := countExcellentPairs([]int{1, 2, 3, 1}, 3)
//	println(a)
//}
