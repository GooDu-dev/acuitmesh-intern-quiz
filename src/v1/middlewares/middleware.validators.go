package middlewares

import (
	"time"

	"github.com/GooDu-Dev/acuitmesh-intern-quiz/src/v1/api/user"
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/src/v1/common"
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/utils"
	customError "github.com/GooDu-Dev/acuitmesh-intern-quiz/utils/error"
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/utils/log"
	"github.com/gin-gonic/gin"
)

func NoValidation(context *gin.Context) {

}

func AuthValidator(context *gin.Context) {
	token, err := context.Cookie("auth-token")
	if err != nil {
		status, res := customError.InvalidHeaderNotAcceptableError.ErrorResponse()
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err.Error())
		context.JSON(status, res)
		context.Abort()
		return
	}
	if common.IsDefaultValueOrNil(token) {
		status, res := customError.MissingRequestError.ErrorResponse()
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), map[string]string{"token": token})
		context.JSON(status, res)
		context.Abort()
		return
	}
	userModel := user.UserModel{}

	userCard, err := userModel.CheckUserToken(token)
	if err != nil {
		status, res := customError.GetErrorResponse(err)
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		context.JSON(status, res)
		context.Abort()
		return
	}
	if common.CompareTimeIsPassed(userCard.ExpiredAt, 60*24) {
		updated_at := time.Now()
		err := userModel.ExpireUserToken(userCard.ID, userCard.Token, updated_at)
		if err != nil {
			status, res := customError.GetErrorResponse(err)
			log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
			context.JSON(status, res)
			context.Abort()
			return
		}
		status, res := customError.UserTokenExpiredError.ErrorResponse()
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), map[string]interface{}{"userCard": userCard})
		context.JSON(status, res)
		context.Abort()
		return
	}
	// go to endpoint
	context.Next()
}
