package main

func isValidSudoku(board [][]byte) bool {
	var occured bool
	var num byte
	for y := 0; y < 9; y++ {
		for num = '1'; num <= '9'; num++ {
			occured = false
			for x := 0; x < 9; x++ {

				if board[y][x] == num {
					if occured {
						return false
					}
					occured = true
				}
			}
		}
	}
	for x := 0; x < 9; x++ {
		for num = '1'; num <= '9'; num++ {
			occured = false
			for y := 0; y < 9; y++ {
				if board[y][x] == num {
					if occured {
						return false
					}
					occured = true

				}
			}
		}
	}
	for xStart := 0; xStart < 9; xStart += 3 {
		for yStart := 0; yStart < 9; yStart += 3 {
			for num = '1'; num <= '9'; num++ {
				occured = false
				for x := xStart; x < xStart+3; x++ {
					for y := yStart; y < yStart+3; y++ {
						if board[y][x] == num {
							if occured {
								return false
							}
							occured = true

						}
					}
				}
			}
		}
	}
	return true
}
