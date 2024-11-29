package user

import (
	"time"

	"github.com/GooDu-Dev/acuitmesh-intern-quiz/src/v1/api/tictactoe"
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/src/v1/common"
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/src/v1/services/auth"
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

		created_at := time.Now()
		token, err := s.Model.CreateUserToken(userCard.ID, created_at)
		if err != nil {
			log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
			return nil, err
		}
		userCard.Token = token.Token
		userCard.ExpiredAt = token.ExpiredAt

	} else if common.CompareTimeIsPassed(userCard.ExpiredAt, 60*24) {
		// token found but expired
		updated_at := time.Now()
		err := s.Model.ExpireUserToken(userCard.ID, userCard.Token, updated_at)
		if err != nil {
			log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
			return nil, err
		}

		created_at := time.Now()
		token, err := s.Model.CreateUserToken(userCard.ID, created_at)
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

	created_at := time.Now()
	if userCard, err = s.Model.CreateUser(request.Email, request.Username, pwd, created_at); err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		return nil, err
	}

	return userCard, nil
}

func (s *UserService) UserCreateInvite(home_token string, away_id int, created_at time.Time) (match *UserInviteStruct, err error) {

	home_id, err := auth.GetService().AuthorizeToken(home_token)
	if err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		return nil, err
	}

	if match, err = s.Model.CheckMatchInvite(*home_id, away_id); err == nil {
		// already have invited
		log.Logging(utils.INFO_LOG, common.GetFunctionWithPackageName(), map[string]interface{}{"match": match})
		return match, nil
	}

	match_token, err := common.GenerateToken(32)
	if err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		return nil, err
	}

	start := time.Now()
	expired := start.Add(24 * time.Hour)

	if match, err = s.Model.UserCreateInvite(*home_id, away_id, match_token, start, expired, created_at); err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		return nil, err
	}

	log.Logging(utils.INFO_LOG, common.GetFunctionWithPackageName(), map[string]interface{}{"match": match})
	return match, nil
}

func (s *UserService) GetUserMatchURL(home_token string, away_id int) (userMatch *UserMatchStruct, err error) {

	home_id, err := auth.GetService().AuthorizeToken(home_token)
	if err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		return nil, err
	}

	var matchInvite *UserInviteStruct
	if matchInvite, err = s.Model.CheckMatchInvite(*home_id, away_id); err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		return nil, err
	}

	userMatch = &UserMatchStruct{
		HomeID: matchInvite.HomeID,
		AwayID: matchInvite.AwayID,
		Token:  matchInvite.Token,
	}

	return userMatch, nil

}

func (s *UserService) AcceptUserMatchToken(match_token string, user_token string) (userMatch *UserMatchStruct, err error) {
	user_id, err := auth.GetService().AuthorizeToken(user_token)
	if err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		return nil, err
	}

	updated_at := time.Now()
	if userMatch, err = s.Model.AcceptUserMatchToken(match_token, *user_id, updated_at); err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		return nil, err
	}

	var _tttmodel tictactoe.TicTacToeModel
	tttmodel := _tttmodel.InitModel()

	created_at := time.Now()
	expired_at := created_at.Add(48 * time.Hour)
	token, err := common.GenerateToken(32)
	if err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		return nil, err
	}

	if err = tttmodel.CreateUserMatch(userMatch.HomeID, userMatch.AwayID, userMatch.InviteID, expired_at, created_at, token); err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		return nil, err
	}

	return userMatch, nil
}
