package user

import (
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/src/v1/common"
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/utils"
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/utils/database"
	customError "github.com/GooDu-Dev/acuitmesh-intern-quiz/utils/error"
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/utils/log"
)

type UserModel struct {
}

func (m *UserModel) InitModel() *UserModel {
	return &UserModel{}
}

func (m *UserModel) GetUsersList(start int) (*[]UsersListResponse, error) {
	var response []UsersListResponse
	result := database.DB.Table("tb_user").
		Select(
			"tb_user.id",
			"tb_user.username",
			"tb_user.win",
			"tb_user.lose",
			"tb_user.draw",
			"tb_user.total",
		).
		Where("tb_user.total <> 0").
		Order("tb_user.total desc, tb_user.win::float / tb_user.total::float").
		Offset(start).
		Limit(50).
		Find(&response)

	if result.Error != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), result.Error)
		return nil, customError.InternalServerError
	}

	return &response, nil
}
