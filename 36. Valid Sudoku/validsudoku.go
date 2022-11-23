package main

import "math"

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

func isValidSudokuImproved(board [][]byte) bool {
	n := byte(len(board))
	level, div := n*n/64, n*n%64
	if div > 0 {
		level += 1
	}
	row_position := make([]uint64, level)
	column_position := make([]uint64, level)
	square_position := make([]uint64, level)
	square := byte(math.Sqrt(float64(n)))

	get := func(b []uint64, n uint) bool {
		pos, rest := n/64, n%64
		return b[pos]&(1<<rest) > 0
	}
	set := func(b []uint64, n uint) {
		pos, rest := n/64, n%64
		b[pos] |= 1 << rest
	}

	for i := byte(0); i < n; i++ {
		for j := byte(0); j < n; j++ {
			value := board[i][j]
			if value == '.' {
				continue
			}
			value -= 48
			idx := uint(i*n + value - 1)
			if get(row_position, idx) {
				return false
			}
			set(row_position, idx)

			idx = uint(j*n + value - 1)
			if get(column_position, idx) {
				return false
			}
			set(column_position, idx)

			idx = uint(((i/square)*square+(j/square))*n + value - 1)
			if get(square_position, idx) {
				return false
			}
			set(square_position, idx)
		}
	}

	return true
}
