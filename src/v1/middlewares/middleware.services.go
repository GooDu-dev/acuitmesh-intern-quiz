package middlewares

import (
	"strings"

	"github.com/GooDu-Dev/acuitmesh-intern-quiz/src/v1/common"
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/utils"
	settings "github.com/GooDu-Dev/acuitmesh-intern-quiz/utils"
	customError "github.com/GooDu-Dev/acuitmesh-intern-quiz/utils/error"
	"github.com/gin-gonic/gin"
)

type ValidatorService struct {
	BasicHeader HeaderRequest
	UserHeader  UserHeaderRequest
}

func (h *HeaderRequest) CheckContentType() (err error) {
	if common.IsDefaultValueOrNil(h.ContentType) {
		return customError.MissingRequestError
	}
	if content_type, err := settings.ContentType.Value(); err == nil {
		if h.ContentType == content_type {
			return nil
		}
		return customError.InvalidHeaderNotAcceptableError
	}
	return err
}

func (h *HeaderRequest) CheckContentCode() (err error) {
	// public key check
	if common.IsDefaultValueOrNil(h.ContentCode) {
		return customError.MissingRequestError
	}
	if content_code, err := settings.ContentCode.Value(); err == nil {
		if h.ContentCode == content_code {
			return nil
		}
		return customError.InvalidHeaderNotAcceptableError
	}
	return nil
}

func (h *HeaderRequest) CheckClientVersion() (version string, err error) {
	// check web version
	lst := strings.Split(h.ClientVersion, ".")
	version = lst[0]
	if v, err := settings.ClientVersion.Value(); err != nil {
		if version == v {
			return version, nil
		}
	}
	return "", customError.InvalidHeaderNotAcceptableError
}

func (h *HeaderRequest) CheckAccessCtrl() (err error) {
	// check user token

	return nil
}

func (h *HeaderRequest) CheckSourceCtrl() (err error) {
	// check api toekn
	return nil
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
