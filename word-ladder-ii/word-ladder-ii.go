package main

import (
	"fmt"
	"time"
)

const wildcard = "?"
const MaxPathLen = 500

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

//func main() {
//	a := findLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"})
//	fmt.Println(a)
//}

func findLadders(beginWord string, endWord string, wordList []string) (ladders [][]string) {
	wordList = append(wordList, beginWord)
	n := len(wordList)
	k := len(beginWord)
	var root = n - 1 // idx of begin word in array of edges
	var dest = n     // idx of end word in array of edges
	edges := make(map[string][]int, n*k)
	for i, word := range wordList {
		for j := range word {
			pattern := word[:j] + wildcard + word[j+1:]
			edges[pattern] = append(edges[pattern], i)
		}
		if word == endWord {
			dest = i
		}
	}

	// BFS
	var pathlen int = -1 // size of a word ladder, equals to the idx of last word in the ladder
	queue := NewQueue(n)
	visited := make([]bool, n)
	depth := make([]int, n)
	froms := make([][]int, n)
	queue.Enqueue(root)
	visited[root] = true
	depth[root] = 0
	for !queue.Empty() {
		node := queue.Dequeue()
		if node == dest {
			pathlen = depth[node]
		}
		// even after finding minimal path len, we continue traversing to gather the
		// reverse paths, but not longer than pathlen
		if pathlen > 0 && depth[node] > pathlen {
			break
		}
		word := wordList[node]
		for i := range word {
			pattern := word[:i] + wildcard + word[i+1:]
			for _, child := range edges[pattern] {
				if !visited[child] {
					visited[child] = true
					depth[child] = depth[node] + 1
					queue.Enqueue(child)
				}
				// Oh, that's a hard one! We do not add child with depth smaller than a current
				// one since it's definitely is not a part of solution. The logic here -- if a
				// child with a smaller depth was a part of solution, then such a path would be
				// more optimal than the current one, so it's either the child with the smaller
				// depth is not a part of solution, or the current iteration isn't. Either way,
				// they should not be interconnected. It's 7am in the morning at Monday, I've
				// gotta get some sleep before work, don't be so harsh on me. Also, since it's a
				// BFS, depth[child] doesn't actually exceed depth[node] by more than 1. The
				// second part is just here to prevent adding the same node number twice, because
				// it's a special bfs, we have an outer loop on pattern, same node can be only
				// added twice while iterating over the same node in the bigger queue loop. Since
				// it's a bfs we do have the outer _for_ iterate over each node just once.
				if depth[child] >= depth[node] && (len(froms[child]) == 0 || froms[child][len(froms[child])-1] != node) {
					froms[child] = append(froms[child], node)
				}
			}
		}

		// BTW _path not found_ case is handled nicely anyway. Without explicit if pathlen==1 return
	}

	// DFS with max depth pathlen to get all the combinations starting from dest
	var paths [][MaxPathLen]int // use arrays to ensure easy copy while recursing
	var DFS func(int, int, [MaxPathLen]int)
	DFS = func(node int, depth int, path [MaxPathLen]int) {
		if depth > pathlen {
			return
		}
		path[depth] = node
		if node == root {
			paths = append(paths, path)
			return
		}
		for _, child := range froms[node] {
			if !isVisited(child, depth, &path) {
				DFS(child, depth+1, path)
			}
		}
	}
	DFS(dest, 0, [MaxPathLen]int{})
	// restore ladders from pathes
	ladders = make([][]string, len(paths))
	for i := range paths {
		ladders[i] = make([]string, pathlen+1)
		for j := 0; j <= pathlen; j++ {
			// reverse the order, since dfs started from the dest
			ladders[i][pathlen-j] = wordList[paths[i][j]]
		}
	}
	return
}
func isVisited(node, depth int, path *[MaxPathLen]int) bool {
	for i := 0; i <= depth; i++ {
		if node == path[i] {
			return true
		}
	}
	return false
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
