package main

import "fmt"

func intersects(a, b segment) bool {
	return (a.direction == b.direction) && (a.secondary == b.secondary) && (a.main2 > b.main1 && a.main1 < b.main1 || b.main2 > a.main1 && b.main1 < a.main2) || (a.direction != b.direction) && a.main1 <= b.secondary && b.secondary <= a.main2 && b.main1 <= a.secondary && a.secondary <= b.main2
}

type Direction bool

const (
	vertical   Direction = false
	horizontal Direction = true
)

type segment struct {
	main1, main2 int
	secondary    int
	direction    Direction
}

const maxDist int = 10e6

func isSelfCrossing(distance []int) bool {
	n := len(distance)
	var path = make([]segment, n)
	var x, y, i int
	for {
		if i >= n {
			break
		}
		path[i] = segment{y, y + distance[i], x, vertical}
		y += distance[i]
		i++
		if i >= n {
			break
		}
		path[i] = segment{x - distance[i], x, y, horizontal}
		x -= distance[i]
		i++
		if i >= n {
			break
		}
		path[i] = segment{y - distance[i], y, x, vertical}
		y -= distance[i]

		i++
		if i >= n {
			break
		}
		path[i] = segment{x, x + distance[i], y, horizontal}
		x += distance[i]
		i++
	}
	if n%4 == 2 || n%4 == 3 {
		path[n-1].main1 -= maxDist
	} else {
		path[n-1].main2 += maxDist
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i-1; j++ {
			if intersects(path[i], path[j]) {
				fmt.Println(path[i], path[j])
				return true
			}
		}
	}
	return false
}

//func Solve() {
//	println(isSelfCrossing(
//		//[]int{1, 1, 1, 1},
//		dirInput,
//	))
//}
