package tictactoe

type TicTacToeService struct {
}

func (s *TicTacToeService) CheckWinner(board []int) (winner int) {
	// Check rows (3 rows in a 3x3 grid)
	for i := 0; i < 3; i++ {
		// Row check (board[i*3], board[i*3+1], board[i*3+2])
		if board[i*3] == board[i*3+1] && board[i*3+1] == board[i*3+2] && board[i*3] != 0 {
			return board[i*3] // PlayerX or PlayerO
		}
	}

	// Check columns (3 columns in a 3x3 grid)
	for i := 0; i < 3; i++ {
		// Column check (board[i], board[i+3], board[i+6])
		if board[i] == board[i+3] && board[i+3] == board[i+6] && board[i] != 0 {
			return board[i] // PlayerX or PlayerO
		}
	}

	// Check diagonals
	// Top-left to bottom-right (board[0], board[4], board[8])
	if board[0] == board[4] && board[4] == board[8] && board[0] != 0 {
		return board[0] // PlayerX or PlayerO
	}

	// Top-right to bottom-left (board[2], board[4], board[6])
	if board[2] == board[4] && board[4] == board[6] && board[2] != 0 {
		return board[2] // PlayerX or PlayerO
	}

	// No winner yet
	return 0
}
