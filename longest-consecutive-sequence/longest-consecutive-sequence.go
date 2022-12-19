package main

import "fmt"

const MinVal = -1_000_000_000
const radix int = 0xff
const MaxCount = 100_000

var counts, zeroes [radix + 1]int
var buffer [MaxCount]int

func CountingSortForRadix(container []int, order int) {
	copy(counts[:], zeroes[:])
	for _, elem := range container {
		counts[(elem>>order)&radix]++
	}
	for i := 1; i <= radix; i++ {
		counts[i] += counts[i-1]
	}
	for i := len(container) - 1; i >= 0; i-- {
		elem := container[i]
		counts[(elem>>order)&radix]--
		buffer[counts[(elem>>order)&radix]] = elem
	}
	copy(container, buffer[:])
}
func RadixSort(container []int) {
	for order := 0x0; order <= 0x3; order++ {
		CountingSortForRadix(container, order*0b1000)
	}
}

func longestConsecutive(nums []int) (ans int) {
	for i, _ := range nums {
		nums[i] += -MinVal
	}
	RadixSort(nums)
	fmt.Println(nums)
	prevNum := MinVal
	count := 0
	for _, num := range nums {
		if num == prevNum {
			continue
		}
		if num != prevNum+1 {
			ans = Max(ans, count)
			count = 0
		}
		prevNum = num
		count++
	}
	return Max(ans, count)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//
//func main() {
//	fmt.Println(a)
//}
