package tictactoe

type TicTacToeService struct {
	Model *TicTacToeModel
}

func (s *TicTacToeService) InitService() *TicTacToeService {
	model := TicTacToeModel{}
	return &TicTacToeService{
		Model: model.InitModel(),
	}
}
