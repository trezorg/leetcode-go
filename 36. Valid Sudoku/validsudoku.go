package main

func isValidSudoku(board [][]byte) bool {
	type position struct {
		x     byte
		y     byte
		value byte
	}
	n := byte(len(board))
	positions := make(map[position]struct{}, 3*n*n)
	pos := position{}

	for i := byte(0); i < n; i++ {
		for j := byte(0); j < n; j++ {
			value := board[i][j]
			if value == '.' {
				continue
			}
			pos.x = i
			pos.y = n
			pos.value = value
			if _, ok := positions[pos]; ok {
				return false
			}
			positions[pos] = struct{}{}
			pos.x = n
			pos.y = j
			if _, ok := positions[pos]; ok {
				return false
			}
			positions[pos] = struct{}{}

			pos.x = i / 3
			pos.y = j / 3
			if _, ok := positions[pos]; ok {
				return false
			}
			positions[pos] = struct{}{}
		}
	}

	return true
}
