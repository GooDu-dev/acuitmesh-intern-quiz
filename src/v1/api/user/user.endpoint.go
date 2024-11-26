package user

import (
	"net/http"
	"strconv"

	"github.com/GooDu-Dev/acuitmesh-intern-quiz/src/v1/common"
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/utils"
	customError "github.com/GooDu-Dev/acuitmesh-intern-quiz/utils/error"
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/utils/log"
	"github.com/gin-gonic/gin"
)

type UserEndpoint struct {
	Service *UserService
}

func NewEndpoint() *UserEndpoint {
	service := UserService{}
	return &UserEndpoint{
		Service: service.InitService(),
	}
}

func (e *UserEndpoint) GetUserHealth(c *gin.Context) {
	log.Logging(utils.INFO_LOG, common.GetFunctionWithPackageName(), "User service health check")
	c.JSON(http.StatusOK, "User service works fine.")
}

func (e *UserEndpoint) GetUsersList(c *gin.Context) {
	// ! DO IN PRODUCTION
	// ! COMMMENT WHEN TEST BCUZ IT'S JUST A QUIZ :)
	// if err := middlewares.CheckBasicHeader(c); err != nil {
	// 	status, res := customError.GetErrorResponse(err)
	// 	log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
	// 	c.JSON(status, res)
	// 	return
	// }

	s := c.Query("s")
	start, err := strconv.Atoi(s)
	if err != nil {
		status, res := customError.InvalidRequestError.ErrorResponse()
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		c.JSON(status, res)
		return
	}

	var response *[]UsersListResponse
	if response, err = e.Service.GetUserList(start); err != nil {
		status, res := customError.GetErrorResponse(err)
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		c.JSON(status, res)
		return
	}

	log.Logging(utils.INFO_LOG, common.GetFunctionWithPackageName(), response)
	c.JSON(http.StatusOK, response)
}

func (e *UserEndpoint) GetUserDashboard(c *gin.Context) {

	_id := c.Query("uid")
	id, err := strconv.Atoi(_id)
	if err != nil {
		status, res := customError.GetErrorResponse(err)
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		c.JSON(status, res)
		return
	}

	var response *UserHistoryResponse
	if response, err = e.Service.GetUserDashboard(id); err != nil {
		status, res := customError.GetErrorResponse(err)
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		c.JSON(status, res)
		return
	}

	c.JSON(http.StatusOK, response)
}
