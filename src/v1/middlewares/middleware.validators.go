package middlewares

import (
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/utils"
	"github.com/gin-gonic/gin"
)

func NoValidation(context *gin.Context) {

}

func CheckBasicHeader(context *gin.Context) (err error) {
	header := HeaderRequest{
		ContentType:   context.GetHeader(utils.CONTENT_TYPE),
		ContentCode:   context.GetHeader(utils.CONTENT_CODE),
		ClientVersion: context.GetHeader(utils.CLIENT_VERSION),
		AccessCtrl:    context.GetHeader(utils.ACCESS_CONTROL),
		SourceCtrl:    context.GetHeader(utils.SOURCE_CONTROL),
	}
	validator := ValidatorService{
		BasicHeader: header,
		UserHeader:  UserHeaderRequest{},
	}

	if err = validator.BasicHeader.CheckContentType(); err != nil {
		return err
	}

	if err = validator.BasicHeader.CheckContentCode(); err != nil {
		return err
	}

	if _, err = validator.BasicHeader.CheckClientVersion(); err != nil {
		return err
	}

	if err = validator.BasicHeader.CheckAccessCtrl(); err != nil {
		return err
	}

	if err = validator.BasicHeader.CheckSourceCtrl(); err != nil {
		return err
	}

	return nil
}
