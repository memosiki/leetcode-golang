package main

//	func Min(x, y int) int {
//		if x < y {
//			return x
//		}
//		return y
//	}
func minimumCardPickup(cards []int) (minDist int) {
	minDist = len(cards) + 1
	var rightmostOccurence = make(map[int]int)
	for i, card := range cards {
		if brother, ok := rightmostOccurence[card]; ok {
			minDist = Min(minDist, i-brother+1)
		}
		rightmostOccurence[card] = i
	}
	if minDist == len(cards)+1 {
		return -1
	}
	return
}

//func main() {
//
//}
