package auth

import (
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/src/v1/common"
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/utils"
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/utils/database"
	customError "github.com/GooDu-Dev/acuitmesh-intern-quiz/utils/error"
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/utils/log"
)

type AuthModel struct {
}

var m *AuthModel

func (m *AuthModel) InitAuthModel() *AuthModel {
	if m != nil {
		return m
	}
	return &AuthModel{}
}

func (m *AuthModel) TokenToID(user_token string) (*int, error) {
	var user_id int
	result := database.DB.Table("tb_user_token").
		Select("tb_user_token.user_id").
		Where("tb_user_token.token = ?", user_token).
		Find(&user_id)

	if result.Error != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), result.Error)
		return nil, customError.UserAuthorizeFailError
	}

	return &user_id, nil
}

func (m *AuthModel) IDToToken(user_id int) (*string, error) {
	var user_token string
	result := database.DB.Table("tb_user_token").
		Select("tb_user_token.user_id").
		Where("tb_user_id = ?", user_id).
		Find(&user_token)

	if result.Error != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), result.Error)
		return nil, customError.UserAuthorizeFailError
	}

	return &user_token, nil
}
