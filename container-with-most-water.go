package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxArea(height []int) (area int) {
	i, j := 0, len(height)-1
	for i < j {
		area = max(area, min(height[i], height[j])*(j-i))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return
}

//func main() {
//	println(maxArea([]int{1,10,1,9,2,10}))
//}
