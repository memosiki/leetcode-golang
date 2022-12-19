package main

//import "github.com/emirpasic/gods/queue/arrayqueue" BRURHRUHRHRRU!

import (
	"container/heap"
)

type Node struct {
	i, j int
}

const MaxVal = 100_000

var scores [MaxVal + 1]int

type any = interface{}

func maxPoints(grid [][]int, queries []int) (ans []int) {
	n := len(grid)
	m := len(grid[0])
	seen := make([][]bool, n)
	for i := range seen {
		seen[i] = make([]bool, m)
	}
	//	var queue = &Queue{}
	//	var laterQueue = &Queue{}
	//	queue.Enqueue(Node{0, 0})
	pq := new(PriorityQueue)
	pq.Push(&Item{Node{0, 0}, grid[0][0], 0})
	var cur = 1
	var elem Node
	var points int
	for pq.Len() > 0 {
		pitem := pq.Pop()
		item := pitem.(*Item)
		elem = item.value
		i, j := elem.i, elem.j
		if grid[elem.i][elem.j] >= cur {
			for ; cur < item.priority; cur++ {
				scores[cur] = points
			}
		}
		seen[i][j] = true
		points++
		if elem.i-1 >= 0 && !seen[i-1][j] {
			pq.Push(&Item{Node{elem.i - 1, elem.j}, grid[i-1][j], 0})
		}
		if elem.j-1 >= 0 && !seen[i][j-1] {
			pq.Push(&Item{Node{elem.i, elem.j - 1}, grid[i][j-1], 0})
		}
		if elem.i+1 < n && !seen[i+1][j] {
			pq.Push(&Item{Node{elem.i + 1, elem.j}, grid[i+1][j], 0})
		}
		if elem.j+1 < m && !seen[i][j+1] {
			pq.Push(&Item{Node{elem.i, elem.j + 1}, grid[i][j+1], 0})
		}
	}
	ans = make([]int, len(queries))
	for i, query := range queries {
		ans[i] = scores[query]
	}
	return
}

//type StructureElement = Node
//type Queue struct {
//	start, end int
//	container  [MaxVal + 1]StructureElement
//}
//
//func (q *Queue) Enqueue(elem StructureElement) {
//	q.container[q.end] = elem
//	q.end++
//}
//func (q *Queue) Dequeue() (elem StructureElement) {
//	elem = q.container[q.start]
//	q.start++
//	return
//}
//func (q *Queue) Empty() bool {
//	return q.end == q.start
//}
//func (q *Queue) Clean() {
//	q.end = 0
//	q.start = 0
//}

type Element = int

func Max(a ...Element) (max Element) {
	max = a[0]
	for _, elem := range a {
		if max < elem {
			max = elem
		}
	}
	return
}
func Min(a ...Element) (min Element) {
	min = a[0]
	for _, elem := range a {
		if min > elem {
			min = elem
		}
	}
	return
}
func Abs(a Element) Element {
	if a < 0 {
		return -a
	}
	return a
}

func Sum(a ...Element) (rolling Element) {
	for _, elem := range a {
		rolling += elem
	}
	return
}

const MOD = 1_000_000_000 + 7

// An Item is something we manage in a priority queue.
type Item struct {
	value    Node // The value of the item; arbitrary.
	priority int  // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value Node, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
