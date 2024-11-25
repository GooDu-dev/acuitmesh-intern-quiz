package middlewares

import (
	"strings"

	"github.com/GooDu-Dev/acuitmesh-intern-quiz/src/v1/common"
	settings "github.com/GooDu-Dev/acuitmesh-intern-quiz/utils"
	customError "github.com/GooDu-Dev/acuitmesh-intern-quiz/utils/error"
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
