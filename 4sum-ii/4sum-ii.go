package main

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) (ans int) {
	n := len(nums1)
	sums := make(map[int]int, n*n)
	for _, elem1 := range nums1 {
		for _, elem2 := range nums2 {
			sums[elem1+elem2]++
		}
	}
	for _, elem3 := range nums3 {
		for _, elem4 := range nums4 {
			ans += sums[-elem3-elem4]
		}
	}
	return
}
