package main

func areSequential(a, b string) bool {
	var diff uint8
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diff++
		}
	}
	return diff == 1
}
func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordList = append(wordList, beginWord)
	n := len(wordList)
	var root = n - 1 // idx of begin word in array ofedges
	var dest = n     // idx of end word in array of edges
	edges := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if areSequential(wordList[i], wordList[j]) {
				edges[i] = append(edges[i], j)
				edges[j] = append(edges[j], i)
			}
		}
		if wordList[i] == endWord {
			dest = i
		}
	}
	queue := NewQueue(n)
	visited := make([]bool, n)
	depth := make([]int, n)
	queue.Enqueue(root)
	visited[root] = true
	depth[root] = 1
	for !queue.Empty() {
		node := queue.Dequeue()
		if node == dest {
			return depth[node]
		}
		for _, child := range edges[node] {
			if !visited[child] {
				visited[child] = true
				depth[child] = depth[node] + 1
				queue.Enqueue(child)
			}
		}
	}
	// not found
	return 0
}

type StructureElement = int

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
