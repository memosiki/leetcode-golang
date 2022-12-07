package main

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func sum(container []int) (runningSum int) {
	for _, elem := range container {
		runningSum += elem
	}
	return
}

func diff(container []int, target int) (difference []int) {
	difference = make([]int, len(container))
	for i, elem := range container {
		difference[i] = abs(target - elem)
	}
	return
}

func findMinMoves(machines []int) int {
	var n int = len(machines)
	dresses := sum(machines)
	if dresses%n != 0 {
		return -1
	}
	var each int = dresses / n
	return sum(diff(machines, each)) / 2

}

//
//func Solve() {
//	println(findMinMoves(
//		[]int{42, 4},
//	))
//}
