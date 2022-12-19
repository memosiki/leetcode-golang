package main

import "fmt"

func smallestStringWithSwaps(s string, pairs [][]int) string {
	n := len(s)
	uf := NewUnionFind(n)

	// split pairs into disjointed sets
	for _, pair := range pairs {
		uf.Union(pair[0], pair[1])
	}

	// gather all disjointed sets aka group
	var groupId int
	groups := make(map[int][]byte)
	for i := 0; i < n; i++ {
		groupId = uf.Find(i)
		groups[groupId] = append(groups[groupId], s[i])
	}
	// sort lexicographically each group
	for _, group := range groups {
		CountingSortInplaceASCII(group)
	}
	fmt.Println(groups)
	builder := make([]byte, n)                 // resulting string, but bytes array
	counters := make(map[int]int, len(groups)) // how many elements from the group were already used
	for i := 0; i < n; i++ {
		groupId = uf.Find(i)
		// use lexicographically smaller symbol from the corresponding group
		builder[i] = groups[groupId][counters[groupId]]
		counters[groupId]++
	}
	return string(builder)
}

func CountingSortInplaceASCII(container []byte) {
	// O(N) complexity

	if len(container) == 1 { // disjointed sets optimization
		return
	}
	const MaxElem = 'z' - 'a'
	// char arithmetics duh
	var counts [MaxElem + 1]int
	for _, elem := range container {
		counts[elem-'a'] += 1
	}
	var i int
	for elem, count := range counts {
		for k := 0; k < count; k++ {
			container[i+k] = byte(elem) + 'a'
		}
		i += count
	}
}

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
