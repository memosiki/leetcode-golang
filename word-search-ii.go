package main

import (
	//	"fmt"
	//	"github.com/emirpasic/gods/sets/hashset"
	//	"github.com/emirpasic/gods/sets"
	"strings"
)

type Trie struct {
	children       map[rune]*Trie // непосредственные сыновья
	parent         *Trie          // родитель
	pathFromParent rune           // по какой букве связан с родителем
	traverse       map[rune]*Trie // переходы, либо в сына либо фоллбек по суффиксной ссылке
	terminals      []int          // список паттернов которые завершаются в данной ноде
	suffLink       *Trie
	isLeaf         bool
}

func (t *Trie) String() string {
	repr := strings.Builder{}

	repr.WriteByte('[')
	for key, val := range t.children {
		repr.WriteRune(key)
		repr.WriteString("->")
		repr.WriteString(val.String())
	}
	if len(t.children) == 0 {
		repr.WriteByte('*')
	}
	repr.WriteByte(']')
	return repr.String()
}
func NewTrie() *Trie {
	links := make(map[rune]*Trie)
	traverse := make(map[rune]*Trie)
	return &Trie{children: links, traverse: traverse}
}

var root *Trie

func (t *Trie) GetSuffLink() *Trie {
	if t.suffLink == nil {
		if t == root || t.parent == root { // по определению суффиксной ссылки
			t.suffLink = root
		} else {
			t.suffLink = t.parent.GetSuffLink().Traverse(t.pathFromParent)
		}
		t.terminals = append(t.terminals, t.suffLink.terminals...)
	}
	return t.suffLink
}
func (t *Trie) Traverse(c rune) *Trie {
	if _, ok := t.traverse[c]; !ok {
		if son, ok := t.children[c]; ok {
			t.traverse[c] = son
		} else if t == root {
			t.traverse[c] = root
		} else {
			t.traverse[c] = t.GetSuffLink().Traverse(c)
		}
	}
	return t.traverse[c]
}
func GenTrie(patterns []string) *Trie {
	root = NewTrie()
	for patternId, pattern := range patterns {
		node := root
		for _, letter := range pattern {
			next, ok := node.children[letter]
			if !ok {
				next = NewTrie()
				next.parent = node
				next.pathFromParent = letter
				node.children[letter] = next
				node.isLeaf = false
			}
			node = next
		}
		node.isLeaf = true
		node.terminals = append(node.terminals, patternId)
	}
	//	BFS(root)
	return root
}

//func BFS(root *Trie) {
//	queue := linkedlistqueue.New()
//	queue.Enqueue(root)
//	for !queue.Empty() {
//		t, _ := queue.Dequeue()
//		node := t.(*Trie)
//		node.GetSuffLink()
//		for _, child := range node.children {
//			queue.Enqueue(child)
//		}
//	}
//}

func SearchPlane(plane [][]byte, patterns []string) (matches []bool) {
	matches = make([]bool, len(patterns))
	trie := GenTrie(patterns)

	var n = len(plane)
	var m = len(plane[0])
	var visited = make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}
	var checkNeighbours func(int, int, *Trie)
	checkNeighbours = func(i, j int, node *Trie) {
		if i < 0 || i >= n || j < 0 || j >= m || visited[i][j] {
			return
		}
		visited[i][j] = true
		node = node.Traverse(rune(plane[i][j]))
		for _, patternId := range node.terminals {
			matches[patternId] = true
		}
		checkNeighbours(i-1, j, node)
		checkNeighbours(i, j-1, node)
		checkNeighbours(i+1, j, node)
		checkNeighbours(i, j+1, node)
		visited[i][j] = false
		return
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			checkNeighbours(i, j, trie)
		}
	}
	return
}

func findWords(
	board [][]byte,
	words []string,
) (matchedPatterns []string) {
	matches := SearchPlane(board, words)
	for i, isFound := range matches {
		if isFound {
			matchedPatterns = append(matchedPatterns, words[i])
		}
	}
	//	trie := GenTrie(words)
	//	fmt.Println(matchedPatterns)
	return
}

//func main() {
//	findWords(
//		[][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}},
//		[]string{"oath", "pea", "eat", "rain"},
//	)
//}
