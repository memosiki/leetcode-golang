package main

func exist(board [][]byte, word string) bool {
	var n = len(board)
	var m = len(board[0])
	var visited = make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}
	var nextLetter func(int, int, int) bool
	nextLetter = func(i, j, curInd int) (ans bool) {
		if i < 0 || i >= n || j < 0 || j >= m || curInd >= len(word) || visited[i][j] {
			return false
		}
		visited[i][j] = true
		if board[i][j] == word[curInd] {
			//            println("checking", i, j, "for", string(word[curInd]))
			if curInd == len(word)-1 {
				return true
			}
			ans = ans || nextLetter(i-1, j, curInd+1)
			ans = ans || nextLetter(i, j-1, curInd+1)
			ans = ans || nextLetter(i+1, j, curInd+1)
			ans = ans || nextLetter(i, j+1, curInd+1)
		}
		visited[i][j] = false
		return
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if nextLetter(i, j, 0) {
				return true
			}
		}
	}
	return false
}

//func main() {
//	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
//    print(exist(board, "ABCCED"))
//}
