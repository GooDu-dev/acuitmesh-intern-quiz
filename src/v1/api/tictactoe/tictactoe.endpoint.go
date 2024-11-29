package tictactoe

import (
	"net/http"
	"strconv"

	"github.com/GooDu-Dev/acuitmesh-intern-quiz/src/v1/common"
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/src/v1/services/tictactoe"
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/utils"
	customError "github.com/GooDu-Dev/acuitmesh-intern-quiz/utils/error"
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/utils/log"
	"github.com/gin-gonic/gin"
)

type TicTacToeEndpoint struct {
	Service *TicTacToeService
}

func NewEndpoint() *TicTacToeEndpoint {
	service := TicTacToeService{}
	return &TicTacToeEndpoint{
		Service: service.InitService(),
	}
}

func (e *TicTacToeEndpoint) GetUserMatchInfo(c *gin.Context) {
	token := c.Query("match")
	if common.IsDefaultValueOrNil(token) {
		status, res := customError.MissingRequestError.ErrorResponse()
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), "match is empty.")
		c.JSON(status, res)
		return
	}

	user_tk, err := c.Cookie("auth-token")
	if err != nil {
		status, res := customError.InvalidRequestError.ErrorResponse()
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		c.JSON(status, res)
		return
	}

	var matchInfo *MatchInfo
	if matchInfo, err = e.Service.GetUserMatchInfo(token, user_tk); err != nil {
		status, res := customError.GetErrorResponse(err)
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		c.JSON(status, res)
		return
	}

	log.Logging(utils.INFO_LOG, common.GetFunctionWithPackageName(), *matchInfo)

	tttService := tictactoe.TicTacToeService{}
	if w := tttService.CheckWinner(matchInfo.Board.Data); w != 0 {
		if w == utils.PLAYER_O {
			c.JSON(http.StatusOK, matchInfo.AwayPlayer)
			return
		} else {
			c.JSON(http.StatusOK, matchInfo.HomePlayer)
			return
		}
	} else {
		c.JSON(http.StatusOK, *matchInfo)
	}

}

func (e *TicTacToeEndpoint) SetXOToBoard(c *gin.Context) {

	_index := c.Query("point")
	index, err := strconv.Atoi(_index)
	if err != nil {
		status, res := customError.InternalServerError.ErrorResponse()
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), "match token is empty.")
		c.JSON(status, res)
		return
	}

	token := c.Query("match")
	if common.IsDefaultValueOrNil(token) {
		status, res := customError.MissingRequestError.ErrorResponse()
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), "match token is empty.")
		c.JSON(status, res)
		return
	}

	mark := c.Query("mark")
	if common.IsDefaultValueOrNil(mark) {
		status, res := customError.MissingRequestError.ErrorResponse()
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), "player mark is empty.")
		c.JSON(status, res)
		return
	}

	var board *BoardInfo
	if board, err = e.Service.SetXOToBoard(token, index, mark); err != nil {
		status, res := customError.MissingRequestError.ErrorResponse()
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), "match token is empty.")
		c.JSON(status, res)
		return
	}

	log.Logging(utils.INFO_LOG, common.GetFunctionWithPackageName(), *board)
	c.JSON(http.StatusOK, *board)

}
