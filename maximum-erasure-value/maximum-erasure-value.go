package main

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func maximumUniqueSubarray(nums []int) (maxSum int) {
	var met = make(map[int]int)
	var sum = make([]int, len(nums)+1) // сумма элементов идущих до указанного
	var left, rollingSum int
	for right, num := range nums {
		if pos, ok := met[num]; ok && pos > left {
			left = pos
		}
		rollingSum += num
		sum[right] = rollingSum
		met[num] = right
		maxSum = Max(maxSum, rollingSum-sum[left])
	}
	return
}

//func main() {
//	println(maximumUniqueSubarray([]int{5, 2, 1, 2, 5, 2, 1, 2, 5}))
//}
