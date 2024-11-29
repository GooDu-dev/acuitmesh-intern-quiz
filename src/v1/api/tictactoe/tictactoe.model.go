package tictactoe

import (
	"time"

	"github.com/GooDu-Dev/acuitmesh-intern-quiz/src/v1/common"
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/utils"
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/utils/database"
	customError "github.com/GooDu-Dev/acuitmesh-intern-quiz/utils/error"
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/utils/log"
)

type TicTacToeModel struct{}

func (m *TicTacToeModel) InitModel() *TicTacToeModel {
	return &TicTacToeModel{}
}

func (m *TicTacToeModel) GetUserMatchInfo(token string, user_id int) (*MatchInfo, error) {

	match := struct {
		HomeID       int       `gorm:"column:home_id"`
		HomeUsername string    `gorm:"column:home_username"`
		AwayID       int       `gorm:"column:away_id"`
		AwayUsername string    `gorm:"column:away_username"`
		InviteID     int       `gorm:"column:invite_id"`
		Token        string    `gorm:"column:token"`
		IsEnd        string    `gorm:"column:is_end"`
		ExpiredAt    time.Time `gorm:"column:expired_at"`
		Board        string    `gorm:"column:board"`
	}{}

	result := database.DB.Table("tb_tictactoe").
		Select(
			"tb_user_home.id as home_id",
			"tb_user_home.username as home_username",
			"tb_user_away.id as away_id",
			"tb_user_away.username as away_username",
			"tb_tictactoe.invite_id",
			"tb_tictactoe.token",
			"tb_tictactoe.is_end",
			"tb_tictactoe.expired_at",
			"tb_tictactoe.board",
		).
		Joins("LEFT JOIN tb_user as tb_user_home ON tb_tictactoe.home_id = tb_user_home.id AND tb_user_home.deleted_at IS NULL").
		Joins("LEFT JOIN tb_user as tb_user_away ON tb_tictactoe.home_id = tb_user_away.id AND tb_user_away.deleted_at IS NULL").
		Where("tb_tictactoe.token = ? AND (tb_tictactoe.home_id = ? OR tb_tictactoe.away_id = ?)", token, user_id, user_id).
		Find(&match)

	if result.Error != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), result.Error)
		return nil, customError.InternalServerError
	}

	data, err := common.StringToIntArray(match.Board)
	if err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		return nil, err
	}

	matchInfo := MatchInfo{
		HomePlayer: UserMatchCard{
			ID:       match.HomeID,
			Username: match.HomeUsername,
		},
		AwayPlayer: UserMatchCard{
			ID:       match.AwayID,
			Username: match.AwayUsername,
		},
		Board: BoardInfo{
			Data:  data,
			Token: match.Token,
		},
		ExpiredAt: match.ExpiredAt,
		InviteID:  match.InviteID,
		IsEnd:     match.IsEnd,
	}

	return &matchInfo, nil
}

func (m *TicTacToeModel) GetUsersInMatch(home_id int, away_id int) (*[]UserMatchCard, error) {

	var usersCard []UserMatchCard

	result := database.DB.Table("tb_user").
		Select("tb_user.id", "tb_user.username").
		Where("tb_user.id = ? OR tb_user.id = ?", home_id, away_id).
		Find(&usersCard)

	if result.Error != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), result.Error)
		return nil, customError.InternalServerError
	}

	// return data ordered as [home_card, away_card]
	if usersCard[0].ID == away_id {
		return &([]UserMatchCard{usersCard[1], usersCard[0]}), nil
	} else {
		return &usersCard, nil
	}
}

func (m *TicTacToeModel) CreateUserMatch(home_id int, away_id int, invite_id int, expired_at time.Time, created_at time.Time, token string) error {
	match := UserMatchCreateStruct{
		HomeID:    home_id,
		AwayID:    away_id,
		InviteID:  invite_id,
		Token:     token,
		ExpiredAt: expired_at,
		IsEnd:     "N",
		Board:     "0,0,0,0,0,0,0,0,0",
		CreatedAt: created_at,
	}

	result := database.DB.Table("tb_tictactoe").
		Select(
			"tb_tictactoe.home_id",
			"tb_tictactoe.away_id",
			"tb_tictactoe.invite_id",
			"tb_tictactoe.token",
			"tb_tictactoe.expired_at",
			"tb_tictactoe.is_end",
			"tb_tictactoe.board",
			"tb_tictactoe.created_at",
		).
		Create(&match)

	if result.Error != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), result.Error)
		return customError.InternalServerError
	}

	return nil
}

func (m *TicTacToeModel) GetBoardFromToken(token string) (*BoardInfo, error) {
	var board BoardInfo
	var boardData struct {
		Data  string `gorm:"column:board"`
		Token string `gorm:"column:token"`
	}

	result := database.DB.Table("tb_tictactoe").
		Select("tb_tictactoe.board", "tb_tictactoe.token").
		Where("tb_tictactoe.token = ?", token).
		Find(&boardData)

	if result.Error != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), result.Error)
		return nil, customError.InternalServerError
	}

	data, err := common.StringToIntArray(boardData.Data)
	if err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		return nil, err
	}
	board = BoardInfo{
		Data:  data,
		Token: boardData.Token,
	}

	return &board, nil
}

func (m *TicTacToeModel) UpdateBoardData(old_token string, new_token string, board_data []int) error {

	result := database.DB.Table("tb_tictactoe").
		Where("tb_tictactoe.token = ?", old_token).
		Updates(map[string]interface{}{
			"token": new_token,
			"board": common.IntArrayToString(board_data),
		})

	if result.Error != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), result.Error)
		return customError.InternalServerError
	}

	return nil
}
