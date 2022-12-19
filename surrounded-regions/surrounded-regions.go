package surrounded_regions

type none = struct{}

func solve(board [][]byte) {
	n := len(board)
	m := len(board[0])

	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, m)
	}
	// DFS to mark all adjacent O-s
	var dfs func(i, j int)

	dfs = func(i, j int) {
		if i < 0 || j < 0 || i >= n || j >= m || visited[i][j] || board[i][j] != 'O' {
			return
		}
		visited[i][j] = true
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}

	// traverse all O-s on the borders, since they are unflippable
	for i := 0; i < m; i++ {
		if board[0][i] == 'O' {
			dfs(0, i)
		}
		if board[n-1][i] == 'O' {
			dfs(n-1, i)
		}
	}
	for i := 0; i < n; i++ {
		if board[i][0] == 'O' {
			dfs(i, 0)
		}
		if board[i][m-1] == 'O' {
			dfs(i, m-1)
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if !visited[i][j] {
				board[i][j] = 'X'
			}
		}
	}
}
