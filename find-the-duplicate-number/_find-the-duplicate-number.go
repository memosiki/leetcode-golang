package main

const MaxVal = 100_000
const flag = 1 << 17

func findDuplicate(nums []int) (ans int) {
	for _, cur := range nums {
		cur = cur &^ flag
		if nums[cur]&flag != 0 {
			ans = cur
			break
		}
		nums[cur] |= flag
	}
	for i := 0; i < len(nums); i++ {
		nums[i] &^= flag
	}
	return
}
