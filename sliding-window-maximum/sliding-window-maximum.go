package main

import "fmt"

const MaxLen = 100_000

func maxSlidingWindow(nums []int, k int) (maxs []int) {
	maxs = make([]int, len(nums)-k+1)
	// deque, contains recently met numbers in decreasing order (maximum first out)
	// (numbers can be outside the window)
	deque := NewDeque(len(nums))

	// counter contains count of numbers in the current window
	// is used to prune numbers outside the window from the queue
	counter := make(Counter, len(nums))

	// organize the window
	for i := 0; i < k-1; i++ {
		for !deque.Empty() && deque.PeekFront() <= nums[i] {
			deque.PopFront()
		}
		for !deque.Empty() && deque.PeekBack() <= nums[i] {
			deque.PopBack()
		}
		deque.PushBack(nums[i])
		counter[nums[i]]++
		//fmt.Println("prepare", i, "queue", deque.Values())
	}
	for i := 0; i < len(nums)-k+1; i++ {
		//fmt.Println(i, "queue", deque.Values())
		cur := nums[k+i-1]
		for !deque.Empty() && (counter[deque.PeekFront()] == 0 || deque.PeekFront() <= cur) {
			deque.PopFront()
		}
		if deque.Empty() {
			maxs[i] = cur
		} else {
			maxs[i] = deque.PeekFront()
		}
		for !deque.Empty() && (counter[deque.PeekBack()] == 0 || deque.PeekBack() <= cur) {
			deque.PopBack()
		}
		deque.PushBack(cur)
		counter[cur]++
		counter[nums[i]]--
	}
	return
}

func main() {
	a := maxSlidingWindow([]int{10, 3, 5, 2, 7, 11, 2}, 3)
	fmt.Println(a)
}

type StructureElement = int
type none = struct{}

// Deque using the array. Any pop/peek operation on the empty deque is undefined
// behaviour. If count of elements inserted overall exceeds dequeCapacity it is
// undefined behaviour, may panic. Good enough for leetcode standards.
type Deque struct {
	start     int // points to the element in the front
	end       int // point on the free space after the element in the back
	container []StructureElement
}

func NewDeque(dequeCapacity int) *Deque {
	container := make([]StructureElement, dequeCapacity*2)
	return &Deque{dequeCapacity, dequeCapacity, container}
}

func (q *Deque) PushBack(elem StructureElement) {
	q.container[q.end] = elem
	q.end++
}
func (q *Deque) PushFront(elem StructureElement) {
	q.start--
	q.container[q.start] = elem
}

func (q *Deque) PopFront() (elem StructureElement) {
	elem = q.container[q.start]
	q.start++
	return
}
func (q *Deque) PopBack() (elem StructureElement) {
	q.end--
	elem = q.container[q.end]
	return
}
func (q *Deque) Empty() bool {
	return q.end == q.start
}
func (q *Deque) PeekFront() (elem StructureElement) {
	return q.container[q.start]
}
func (q *Deque) PeekBack() (elem StructureElement) {
	return q.container[q.end-1]
}
func (q *Deque) Values() []StructureElement {
	return q.container[q.start:q.end]
}

// Counter is a limited port of collections.Counter from Python
type Counter map[StructureElement]int
