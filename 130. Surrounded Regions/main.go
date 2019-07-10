package main

import "log"

// Given a 2D board containing 'X' and 'O' (the letter O), capture all regions surrounded by 'X'.
//
// A region is captured by flipping all 'O's into 'X's in that surrounded region.
func main() {
	ex1 := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	solve(ex1)
	log.Println(ex1)
	/*
		X X X X
		X X X X
		X X X X
		X O X X
	*/

	ex2 := [][]byte{
		{'X', 'X', 'O'},
		{'X', 'O', 'X'},
		{'O', 'X', 'O'},
	}
	solve(ex2)
	log.Println(ex2)
	/*
		X X O
		X X X
		O X O
	*/
}

func solve(board [][]byte) {
	m := len(board)
	if m == 0 {
		return
	}

	n := len(board[0])
	if n == 0 {
		return
	}

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if isBorder(r, c, m, n) {
				dfs(board, r, c, m, n)
			}
		}
	}

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			switch board[r][c] {
			case 'T':
				board[r][c] = 'O'
			case 'O':
				board[r][c] = 'X'
			}
		}
	}
}

func isBorder(r, c, m, n int) bool {
	if r == 0 || r == m-1 || c == 0 || c == n-1 {
		return true
	}
	return false
}

func dfs(board [][]byte, r, c, m, n int) {
	if !sanityCheck(r, c, m, n) {
		return
	}

	switch board[r][c] {
	case 'X', 'T':
		return
	}

	board[r][c] = 'T'
	dfs(board, r-1, c, m, n)
	dfs(board, r+1, c, m, n)
	dfs(board, r, c-1, m, n)
	dfs(board, r, c+1, m, n)
}

func sanityCheck(r, c, m, n int) bool {
	if r < 0 || r > m-1 || c < 0 || c > n-1 {
		return false
	}
	return true
}
