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
		return nil, customError.BadRequestError
	}

	if response, err = s.Model.GetUsersList(start); err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		return nil, err
	}

	return response, nil
}

func (s *UserService) GetUserStatistic(id int) (response *UserStatistic, err error) {

	if id < 0 {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), "id < 0")
		return nil, customError.BadRequestError
	}

	if response, err = s.Model.GetUserStatistic(id); err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		return nil, err
	}

	return response, nil
}

func (s *UserService) GetUserHistoryMatch(id int) (response *[]HistoryMath, err error) {

	if id < 0 {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), "id < 0")
		return nil, customError.BadRequestError
	}

	if response, err = s.Model.GetUserHistoryMatch(id); err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), "id < 0")
		return nil, err
	}

	return response, nil
}

func (s *UserService) GetUserDashboard(id int) (response *UserHistoryResponse, err error) {

	if id < 0 {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		return nil, customError.BadRequestError
	}

	var stat *UserStatistic
	var history *[]HistoryMath
	if stat, err = s.Model.GetUserStatistic(id); err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		return nil, err
	}
	if history, err = s.Model.GetUserHistoryMatch(id); err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		return nil, err
	}

	response = &UserHistoryResponse{
		Win:   stat.Win,
		Lose:  stat.Lose,
		Draw:  stat.Draw,
		Total: stat.Total,
		Match: *history,
	}

	return response, nil

}

func (s *UserService) LoginUser(request UserLoginRequest) (userCard *UserCard, err error) {

	if !common.IsValidEmail(request.Email) {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), map[string]string{"email": request.Email})
		return nil, customError.BadRequestError
	}

	pwd, err := common.UserEncrypt(request.Password)
	if err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		return nil, customError.InternalServerError
	}

	userCard, err = s.Model.LoginUser(request.Email, pwd)
	if err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		return nil, err
	}

	if common.IsDefaultValueOrNil(userCard.Token) {
		// no token found
		token, err := s.Model.CreateUserToken(userCard.ID)
		if err != nil {
			log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
			return nil, err
		}
		userCard.Token = token.Token
		userCard.ExpiredAt = token.ExpiredAt

	} else if common.CompareTimeIsPassed(userCard.ExpiredAt, 60*24) {
		// token found but expired
		err := s.Model.ExpireUserToken(userCard.ID, userCard.Token)
		if err != nil {
			log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
			return nil, err
		}
		token, err := s.Model.CreateUserToken(userCard.ID)
		if err != nil {
			log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
			return nil, err
		}
		userCard.Token = token.Token
		userCard.ExpiredAt = token.ExpiredAt
	}

	return userCard, nil

}

func (s *UserService) CreateUser(request UserRegisterRequest) (userCard *UserCard, err error) {

	if !common.IsValidEmail(request.Email) {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), map[string]string{"email": request.Email})
		return nil, customError.BadRequestError
	}

	pwd, err := common.UserEncrypt(request.Pwd)
	if err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		return nil, customError.InternalServerError
	}

	if userCard, err = s.Model.CreateUser(request.Email, request.Username, pwd); err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		return nil, err
	}

	return userCard, nil
}
