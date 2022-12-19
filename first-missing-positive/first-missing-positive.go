package main

func firstMissingPositive(nums []int) (ans int) {
	n := len(nums)
	maxNum := n + 0xdead
	cur := 0
	for i := 0; i < n; i++ {
		cur = nums[i]
		for 0 < cur && cur <= n {
			cur, nums[cur-1] = nums[cur-1], maxNum
		}
	}
	for ans = 0; ans < n && nums[ans] == maxNum; ans++ {
	}
	ans++
	return
}
