package main

import "log"

// On an 8 x 8 chessboard, there is one white rook.  There also may be empty squares,
// white bishops, and black pawns.  These are given as characters 'R', '.', 'B', and 'p'
// respectively. Uppercase characters represent white pieces, and lowercase characters
// represent black pieces.
//
// The rook moves as in the rules of Chess: it chooses one of four cardinal directions
// (north, east, west, and south), then moves in that direction until it chooses to stop,
// reaches the edge of the board, or captures an opposite colored pawn by moving to the
// same square it occupies.  Also, rooks cannot move into the same square as other friendly bishops.
//
// Return the number of pawns the rook can capture in one move.
func main() {
	ex1 := [][]byte{
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'.', '.', '.', 'R', '.', '.', '.', 'p'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
	}
	log.Println(numRookCaptures(ex1)) // 3

	ex2 := [][]byte{
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', 'p', 'p', 'p', 'p', 'p', '.', '.'},
		{'.', 'p', 'p', 'B', 'p', 'p', '.', '.'},
		{'.', 'p', 'B', 'R', 'B', 'p', '.', '.'},
		{'.', 'p', 'p', 'B', 'p', 'p', '.', '.'},
		{'.', 'p', 'p', 'p', 'p', 'p', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
	}
	log.Println(numRookCaptures(ex2)) // 0

	ex3 := [][]byte{
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'p', 'p', '.', 'R', '.', 'p', 'B', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'B', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
	}
	log.Println(numRookCaptures(ex3)) // 3
}

func numRookCaptures(board [][]byte) int {
	rookR, rookC := 0, 0
	for r := 0; r < 8; r++ {
		rookFinded := false
		for c := 0; c < 8; c++ {
			if board[r][c] == 'R' {
				rookFinded = true
				rookR, rookC = r, c
				break
			}
		}
		if rookFinded {
			break
		}
	}

	pawns := 0
	// see left
	if rookC > 0 {
		for c := rookC - 1; c >= 0; c-- {
			if board[rookR][c] == 'B' {
				break
			} else if board[rookR][c] == 'p' {
				pawns++
				break
			}
		}
	}

	// see right
	if rookC < 7 {
		for c := rookC + 1; c < 8; c++ {
			if board[rookR][c] == 'B' {
				break
			} else if board[rookR][c] == 'p' {
				pawns++
				break
			}
		}
	}

	// see top
	if rookR > 0 {
		for r := rookR - 1; r >= 0; r-- {
			if board[r][rookC] == 'B' {
				break
			} else if board[r][rookC] == 'p' {
				pawns++
				break
			}
		}
	}

	// see bottom
	if rookR < 7 {
		for r := rookR + 1; r < 8; r++ {
			if board[r][rookC] == 'B' {
				break
			} else if board[r][rookC] == 'p' {
				pawns++
				break
			}
		}
	}

	return pawns
}
