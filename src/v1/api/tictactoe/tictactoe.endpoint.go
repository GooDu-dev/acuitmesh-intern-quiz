package tictactoe

import (
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/utils"
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/utils/common"
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

func (e *TicTacToeEndpoint) SetXOToBoard(c *gin.Context) {

	var request BoardInfo
	if err := c.BindJSON(&request); err != nil {
		status, res := customError.InvalidRequestError.ErrorResponse()
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		c.JSON(status, res)
		return
	}

}
