package main

func gameOfLife(board [][]int) {
	n := len(board)
	m := len(board[0])
	//prevBoard := make([][]int, n)
	//for i := 0; i < n; i++ {
	//	prevBoard[i] = make([]int, m)
	//	copy(prevBoard[i], board[i])
	//}
	var neighbours = func(i, j int) (count int) {
		if 0 <= i-1 {
			if 0 <= j-1 {
				count += board[i-1][j-1] & 1
			}
			count += board[i-1][j] & 1
			if m > j+1 {
				count += board[i-1][j+1] & 1
			}
		}
		if n > i+1 {
			if 0 <= j-1 {
				count += board[i+1][j-1] & 1
			}
			count += board[i+1][j] & 1
			if m > j+1 {
				count += board[i+1][j+1] & 1
			}
		}
		if 0 <= j-1 {
			count += board[i][j-1] & 1
		}
		if m > j+1 {
			count += board[i][j+1] & 1
		}
		return
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			switch neighbours(i, j) {
			case 2:
				// same state
				board[i][j] |= board[i][j] << 1
			case 3:
				// alive
				board[i][j] |= 2
			case 0, 1, 4, 5, 6, 7, 8:
				// death
				// noop
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			board[i][j] >>= 1
		}
	}
}
