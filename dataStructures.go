package main

// Type aliases used in place of generics since leetcode runs Go 1.17.6
type any = interface{} // actually means any comparable
type StructureElement = any
type defaultdict = map[any]int // handy reminder that maps in golang actually operate as defaultdict

/* ----------------------------Counter-----------------------------------*/

// Counter is a limited port of collections.Counter from Python
type Counter map[StructureElement]int

// Intersection returns counter that is an intersection of operands element-wise
func (counter Counter) Intersection(other Counter) (counts Counter) {
	counts = make(map[StructureElement]int)
	for key, value := range counter {
		if other[key] > 0 {
			counts[key] = Min2(value, other[key])
		}
	}
	return
}
func (counter Counter) IntersectionInplace(other Counter) {
	for key, value := range counter {
		if other[key] > 0 {
			counter[key] = Min2(value, other[key])
		} else {
			delete(counter, key)
		}
	}
}
func (counter Counter) Union(other Counter) (counts Counter) {
	counts = make(map[StructureElement]int)
	for key, value := range counter {
		counts[key] = value
	}
	for key, value := range other {
		counts[key] += value
	}
	return
}
func (counter Counter) UnionInplace(other Counter) {
	for key, value := range other {
		counter[key] += value
	}
}

// GetAny returns any element
func (counter Counter) GetAny() (key StructureElement, value int, ok bool) {
	for k, val := range counter {
		return k, val, true
	}
	return 0, 0, false
}

// IsUnique returns true if all elements in counter are unique
func (counter Counter) IsUnique() bool {
	for _, count := range counter {
		if count > 1 {
			return false
		}
	}
	return true
}
func NewCounter(container []StructureElement) Counter {
	counts := make(map[StructureElement]int, len(container))
	for _, element := range container {
		counts[element] += 1
	}
	return counts
}
func (counter Counter) Copy() Counter {
	counts := make(map[StructureElement]int, len(counter))
	for k, v := range counter {
		counts[k] = v
	}
	return counts
}

/*----------------Stack-----------------*/

type Stack []StructureElement

func (s Stack) Push(v StructureElement) Stack {
	return append(s, v)
}
func (s Stack) Pop() (Stack, StructureElement) {
	return s[:len(s)-1], s[len(s)-1]
}
func (s Stack) Empty() bool {
	return 0 < len(s)
}

/*------------------Set-------------------*/

type Set map[StructureElement]none
type none = struct{}

func (s Set) Add(val StructureElement) {
	s[val] = none{}
}
func (s Set) In(val StructureElement) bool {
	_, ok := s[val]
	return ok
}
func (s Set) Remove(val StructureElement) {
	delete(s, val)
}

/*--------------------Queue----------------------*/

// Queue is an array backed queue, panics after total queueCapacity elements enqueued.
// Good enough for leetcode standards
type Queue struct {
	start, end int
	container  []StructureElement
}

func NewQueue(queueCapacity int) *Queue {
	container := make([]StructureElement, queueCapacity)
	return &Queue{container: container}
}

func (q *Queue) Enqueue(elem StructureElement) {
	q.container[q.end] = elem
	q.end++
}
func (q *Queue) Dequeue() (elem StructureElement) {
	elem = q.container[q.start]
	q.start++
	return
}
func (q *Queue) Peek() (elem StructureElement) {
	return q.container[q.start]
}
func (q *Queue) Values() []StructureElement {
	return q.container[q.start:q.end]
}
func (q *Queue) Empty() bool {
	return q.end == q.start
}

/*----------------------------Deque-------------------------------*/

// Deque using the array. Any pop/peek operation on the empty deque is undefined
// behaviour. If count of elements inserted overall exceeds dequeCapacity it is
// undefined behaviour, may panic. Good enough for leetcode standards.
type Deque struct {
	start     int // points to the element in the front
	end       int // point to the free space after the last element
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

/*--------------------UnionFind--------------------------*/

// UnionFind is a disjoint-set data structure.
//
// Implementation uses path compression and rank heuristic.
// T(2N)->O(N) space, can be lowered to T(N) by using RandomUnion heuristic
// in place of rank.
// Find, Union, Connected -- O(α(N)), where α is inverse Ackermann function.
type UnionFind struct {
	parent []int
	rank   []int
}

func NewUnionFind(size int) *UnionFind {
	parent := make([]int, size)
	rank := make([]int, size)
	for i := range parent {
		parent[i] = i
		rank[i] = 1
	}
	return &UnionFind{parent, rank}
}
func (uf *UnionFind) Find(node int) int {
	// two approaches are implemented
	const RecursiveApproach = false
	if RecursiveApproach {
		if uf.parent[node] == node {
			return node
		}
		uf.parent[node] = uf.Find(uf.parent[node])
		return uf.parent[node]
	} else {
		root := uf.parent[node]
		for root != uf.parent[root] {
			root = uf.parent[root]
		}
		for node != root {
			node, uf.parent[node] = uf.parent[node], root
		}
		return root
	}
}
func (uf *UnionFind) Union(x, y int) {
	x = uf.Find(x)
	y = uf.Find(y)
	if x != y {
		if uf.rank[x] < uf.rank[y] {
			uf.parent[x] = y
		} else if uf.rank[x] > uf.rank[y] {
			uf.parent[y] = x
		} else {
			uf.parent[y] = x
			uf.rank[x] += 1
		}
	}
}
func (uf *UnionFind) Connected(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}
