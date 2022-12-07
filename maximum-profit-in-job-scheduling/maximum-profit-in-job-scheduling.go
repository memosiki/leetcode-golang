package main

import (
	"github.com/emirpasic/gods/maps/treemap"
	"sort"
)

type Task struct {
	start, end, profit int
}

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	n := len(startTime)
	tasks := make([]Task, n)
	for i := range tasks {
		tasks[i] = Task{startTime[i], endTime[i], profit[i]}
	}
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].end < tasks[j].end
	})
	tm := treemap.NewWithIntComparator()

	tm.Put(0, 0)
	for _, task := range tasks {
		_, valFloor := tm.Floor(task.start)
		_, valMax := tm.Max()
		profitsBefore, profitsSoFar := valFloor.(int), valMax.(int)
		if profitsBefore+task.profit > profitsSoFar {
			tm.Put(task.end, profitsBefore+task.profit)
		}
	}
	_, totalProfit := tm.Max()
	return totalProfit.(int)
}

func main() {
	a := jobScheduling(
		[]int{1, 4, 8, 16, 100, 124, 1040, 2089, 19230},
		[]int{100, 4, 8, 16, 100, 124, 1040, 2089, 19230},
		[]int{1, 4, 8, 16, 100, 124, 1040, 2089, 19230},
	)
	println(a)
}
