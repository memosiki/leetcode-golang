package main

func findDuplicate(nums []int) int {
	hare, tort := 0, 0
	for {
		hare = nums[hare]
		hare = nums[hare]
		tort = nums[tort]
		if hare == tort {
			break
		}
	}
	tort = 0
	for {
		hare = nums[hare]
		tort = nums[tort]
		if hare == tort {
			return hare
		}
	}
}
