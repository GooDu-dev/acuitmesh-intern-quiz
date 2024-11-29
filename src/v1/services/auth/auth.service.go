package auth

import (
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/src/v1/common"
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/utils"
	customError "github.com/GooDu-Dev/acuitmesh-intern-quiz/utils/error"
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/utils/log"
)

type AuthService struct {
	Model *AuthModel
}

var s *AuthService

func GetService() *AuthService {
	if s != nil {
		return s
	}
	s = &AuthService{}
	return s.InitService()
}

func (s *AuthService) InitService() *AuthService {
	model := AuthModel{}
	return &AuthService{
		Model: model.InitAuthModel(),
	}
}

func (s *AuthService) AuthorizeToken(user_token string) (user_id *int, err error) {
	if common.IsDefaultValueOrNil(user_token) {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), "User token is empty.")
		return nil, customError.UserAuthorizeFailError
	}

	if user_id, err = s.Model.TokenToID(user_token); err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), "Cannot authorize user token.")
		return nil, customError.UserAuthorizeFailError
	}

	return user_id, nil
}

func (s *AuthService) DeAuthorizeToken(user_id int) (user_token *string, err error) {
	if common.IsDefaultValueOrNil(user_id) {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), "User token is empty.")
		return nil, customError.UserAuthorizeFailError
	}

	if user_token, err = s.Model.IDToToken(user_id); err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), "Cannot authorize user token.")
		return nil, customError.UserAuthorizeFailError
	}

	return user_token, nil
}
