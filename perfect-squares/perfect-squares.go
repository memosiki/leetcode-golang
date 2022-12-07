package main

import "math"

func numSquares(n int) int {
	var perfectSquares []int
	var optimal = make([]int, n+1)
	for i := range optimal {
		optimal[i] = math.MaxInt32
	}
	for i := 1; i*i <= n; i++ {
		perfectSquares = append(perfectSquares, i*i)
		optimal[i*i] = 1
	}
	for i := range optimal {
		for _, square := range perfectSquares {
			if i+square > n {
				break
			}
			optimal[i+square] = min(optimal[i+square], optimal[i]+1)
		}
	}
	return optimal[n]
}

//
//func main() {
//    println(numSquares())
//}
