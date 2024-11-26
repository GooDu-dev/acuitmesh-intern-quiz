package user

import (
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/src/v1/common"
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/utils"
	customError "github.com/GooDu-Dev/acuitmesh-intern-quiz/utils/error"
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/utils/log"
)

type UserService struct {
	Model *UserModel
}

func (s *UserService) InitService() *UserService {
	model := UserModel{}
	return &UserService{
		Model: model.InitModel(),
	}
}

func (s *UserService) GetUserList(start int) (response *[]UsersListResponse, err error) {

	if start < 0 {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), map[string]int{"start": start})
		return nil, customError.InvalidRequestError
	}

	if response, err = s.Model.GetUsersList(start); err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		return nil, err
	}

	return response, nil
}
