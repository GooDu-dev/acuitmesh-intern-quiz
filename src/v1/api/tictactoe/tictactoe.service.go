package tictactoe

import (
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/src/v1/common"
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/src/v1/services/auth"
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/utils"
	customError "github.com/GooDu-Dev/acuitmesh-intern-quiz/utils/error"
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/utils/log"
)

type TicTacToeService struct {
	Model *TicTacToeModel
}

func (s *TicTacToeService) InitService() *TicTacToeService {
	model := TicTacToeModel{}
	return &TicTacToeService{
		Model: model.InitModel(),
	}
}

func (s *TicTacToeService) GetUserMatchInfo(token string, user_tk string) (matchInfo *MatchInfo, err error) {
	user_id, err := auth.GetService().AuthorizeToken(user_tk)
	if err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		return nil, err
	}

	if matchInfo, err = s.Model.GetUserMatchInfo(token, *user_id); err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		return nil, err
	}

	return matchInfo, nil
}

func (s *TicTacToeService) SetXOToBoard(token string, index int, mark string) (board *BoardInfo, err error) {

	if board, err = s.Model.GetBoardFromToken(token); err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		return nil, err
	}

	if index > 9 || index < 0 {
		// index out of length
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), "index out of length")
		return nil, customError.BadRequestError
	}

	if board.Data[index] != 0 {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), map[string]interface{}{"board": board, "index": index})
		return nil, customError.BadRequestError
	}

	if mark == "x" {
		board.Data[index] = utils.PLAYER_X
	} else if mark == "o" {
		board.Data[index] = utils.PLAYER_O
	} else {
		// invalid input
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		return nil, customError.BadRequestError
	}

	new_token, err := common.GenerateToken(32)
	if err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		return nil, err
	}

	if err := s.Model.UpdateBoardData(board.Token, new_token, board.Data); err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		return nil, customError.InternalServerError
	}

	board = &BoardInfo{
		Data:  board.Data,
		Token: new_token,
	}

	return board, nil

}
