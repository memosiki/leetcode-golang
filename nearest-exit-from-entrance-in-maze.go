package main

type Node struct {
	i, j  int
	depth int
}

func (n *Node) Children() (children []*Node) {
	if n.i != 0 && isEmpty(maze[n.i-1][n.j]) && !visited[n.i-1][n.j] {
		children = append(children, NewNode(n.i-1, n.j, n.depth+1))
	}
	if n.i != N-1 && isEmpty(maze[n.i+1][n.j]) && !visited[n.i+1][n.j] {
		children = append(children, NewNode(n.i+1, n.j, n.depth+1))
	}
	if n.j != 0 && isEmpty(maze[n.i][n.j-1]) && !visited[n.i][n.j-1] {
		children = append(children, NewNode(n.i, n.j-1, n.depth+1))
	}
	if n.j != M-1 && isEmpty(maze[n.i][n.j+1]) && !visited[n.i][n.j+1] {
		children = append(children, NewNode(n.i, n.j+1, n.depth+1))
	}
	return
}
func (n *Node) IsExit() bool {
	return n.depth > 0 && (n.i == 0 || n.j == 0 || n.i == N-1 || n.j == M-1)
}
func isEmpty(char byte) bool {
	return char == '.'
}
func NewNode(i, j, depth int) *Node {
	visited[i][j] = true
	return &Node{i: i, j: j, depth: depth}
}

var (
	maze    [][]byte
	M       int
	N       int
	visited [][]bool
)

func setMaze(mazeInit [][]byte) {
	maze = mazeInit
	N = len(maze)
	M = len(maze[0])
	visited = make([][]bool, N)
	for i := range visited {
		visited[i] = make([]bool, M)
	}
}

func nearestExit(maze [][]byte, entrance []int) (pathlen int) {
	setMaze(maze)
	var node *Node = NewNode(entrance[0], entrance[1], 0)
	queue := make(chan *Node, N*M+1)
	queue <- node
	for {
		select {
		case node = <-queue:
			if node.IsExit() {
				return node.depth
			}
			for _, node := range node.Children() {
				queue <- node
			}
		default:
			return -1
		}
	}
}

//
//func Solve() {
//    fmt.Println(nearestExit(
//		mazeInput,
//        []int{42, 4},
//	))
//}
