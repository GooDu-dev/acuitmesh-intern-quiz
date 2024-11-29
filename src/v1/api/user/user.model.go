package user

import (
	"time"

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

func (m *UserModel) GetUserStatistic(id int) (*UserStatistic, error) {
	var response UserStatistic

	result := database.DB.Table("tb_user").
		Select(
			"tb_user.win",
			"tb_user.lose",
			"tb_user.draw",
			"tb_user.total",
		).
		Where("tb_user.id = ?", id).
		First(&response)

	if result.Error != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), result.Error)
		return nil, customError.InternalServerError
	}

	return &response, nil
}

func (m *UserModel) GetUserHistoryMatch(user_id int) (*[]HistoryMath, error) {

	var query []struct {
		ID           int    `gorm:"column:id"`
		HomeID       int    `gorm:"column:home_id"`
		HomeUsername string `gorm:"column:home_name"`
		AwayID       int    `gorm:"column:away_id"`
		AwayUsername string `gorm:"column:away_name"`
		Winner       string `gorm:"column:winner"`
	}

	result := database.DB.Table("tb_history").
		Select(
			"tb_history.id",
			"tb_history.winner",
			"tb_history.invite_id",
			"tb_history.home_id",
			"home.username as home_name",
			"tb_history.away_id",
			"away.username as away_name",
		).
		Joins("INNER JOIN tb_user AS home ON tb_history.home_id = home.id").
		Joins("INNER JOIN tb_user AS away ON tb_history.away_id = away.id").
		Where("tb_history.home_id = ? OR tb_history.away_id = ?", user_id, user_id).
		Order("tb_history.id desc").
		Limit(20).
		Find(&query)

	if result.Error != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), result.Error)
		return nil, customError.InternalServerError
	}

	var response []HistoryMath
	for _, q := range query {
		response = append(response, HistoryMath{
			ID: q.ID,
			Home: UserCard{
				ID:       q.HomeID,
				Username: q.HomeUsername,
			},
			Away: UserCard{
				ID:       q.AwayID,
				Username: q.AwayUsername,
			},
			Winner: q.Winner,
		})
	}

	return &response, nil

}

func (m *UserModel) LoginUser(email string, pwd string) (*UserCard, error) {

	var user UserCard

	result := database.DB.Table("tb_user").
		Select(
			"tb_user.id",
			"tb_user.username",
			"tb_user_token.token",
			"tb_user_token.on_used",
			"tb_user_token.expired_at",
		).
		Joins("LEFT JOIN tb_user_token ON tb_user.id = tb_user_token.user_id AND tb_user_token.on_used = 'A'").
		Where("tb_user.mail = ? AND tb_user.pwd = ?", email, pwd).
		Find(&user)

	if result.Error != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), result.Error)
		return nil, customError.InternalServerError
	}
	if common.IsDefaultValueOrNil(user.ID) {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), user)
		return nil, customError.UserNotFoundError
	}

	return &user, nil
}

func (m *UserModel) CreateUser(email string, username string, pwd string, created_at time.Time) (*UserCard, error) {

	user := UserCreateStruct{
		Username:  username,
		Email:     email,
		Pwd:       pwd,
		CreatedAt: created_at,
	}

	result := database.DB.Table("tb_user").
		Select("username", "mail", "pwd", "created_at").
		Create(&user)

	if result.Error != nil {
		if common.IsPostgresqlDataDup(result.Error) {
			log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), "data dupplicated")
			return nil, customError.UserAccountDupplicatedError
		}
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), result.Error)
		return nil, customError.InternalServerError
	}

	userCard := UserCard{
		ID:       int(user.ID),
		Username: user.Username,
	}

	return &userCard, nil

}

func (m *UserModel) CheckUserToken(token string) (*UserCard, error) {

	var userCard UserCard

	result := database.DB.Table("tb_user").
		Select(
			"tb_user.id",
			"tb_user.username",
			"tb_user_token.token",
			"tb_user_token.expired_at",
		).
		Joins("LEFT JOIN tb_user_token ON tb_user.id = tb_user_token.user_id AND tb_user_token.on_used = 'A'").
		Where("tb_user_token.token = ?", token).
		First(&userCard)

	if result.Error != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), result.Error)
		return nil, customError.InternalServerError
	}

	return &userCard, nil
}

func (m *UserModel) CreateUserToken(user_id int, created_at time.Time) (*UserTokenStruct, error) {

	tk, err := common.GenerateToken(32)
	if err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		return nil, err
	}

	token := UserTokenStruct{
		UserID:    user_id,
		ExpiredAt: time.Now().Add(24 * time.Hour),
		Token:     tk,
		OnUsed:    "A",
		CreatedAt: created_at,
	}

	result := database.DB.Table("tb_user_token").
		Select(
			"tb_user_token.user_id",
			"tb_user_token.expired_at",
			"tb_user_token.token",
			"tb_user_token.on_used",
			"tb_user_token.created_at",
		).
		Create(&token)

	if result.Error != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), result.Error)
		return nil, customError.InternalServerError
	}

	return &token, nil
}

func (m *UserModel) ExpireUserToken(user_id int, token string, updated_at time.Time) error {

	result := database.DB.Table("tb_user_token").
		Where("tb_user_token.user_id = ? AND tb_user_token.token = ?", user_id, token).
		Updates(map[string]interface{}{
			"on_used":    "N",
			"updated_at": updated_at,
		})

	if result.Error != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), result.Error)
		return customError.InternalServerError
	}

	return nil
}

func (m *UserModel) UserCreateInvite(home_id int, away_id int, match_token string, start time.Time, expired time.Time, created_at time.Time) (*UserInviteStruct, error) {

	invitation := UserInviteStruct{
		HomeID:      home_id,
		AwayID:      away_id,
		StartTime:   start,
		ExpiredTime: expired,
		IsAccept:    "N",
		Token:       match_token,
		CreatedAt:   &start,
	}

	result := database.DB.Table("tb_invitation").
		Select(
			"tb_invitation.start_time_stamp",
			"tb_invitation.expired_time_stamp",
			"tb_invitation.is_accept",
			"tb_invitation.home_id",
			"tb_invitation.away_id",
			"tb_invitation.token",
			"tb_invitation.created_at",
		).
		Create(&invitation)

	if result.Error != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), result.Error)
		return nil, customError.InternalServerError
	}

	return &invitation, nil
}

func (m *UserModel) CheckMatchInvite(home_id int, away_id int) (*UserInviteStruct, error) {

	var match UserInviteStruct

	result := database.DB.Table("tb_invitation").
		Select(
			"tb_invitation.start_time_stamp",
			"tb_invitation.expired_time_stamp",
			"tb_invitation.is_accept",
			"tb_invitation.home_id",
			"tb_invitation.away_id",
			"tb_invitation.token",
			"tb_invitation.created_at",
		).
		Where("tb_invitation.home_id = ? AND tb_invitation.away_id = ? AND tb_invitation.is_accept = 'N'", home_id, away_id).
		Find(&match)

	if result.Error != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), result.Error)
		return nil, customError.InternalServerError
	}
	if err := common.DeepIsDefaultValueOrNil(match); err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		return nil, err
	}

	return &match, nil
}

func (m *UserModel) GetUserMatch(token string, home_id int) (*UserMatchStruct, error) {

	var userMatch UserMatchStruct

	result := database.DB.Table("tb_invitation").
		Select(
			"tb_invitation.home_id",
			"tb_invitation.away_id",
			"tb_invitation.token",
		).
		Where("tb_invitation.token = ? AND tb_invitation.home_id = ?", token, home_id).
		Find(&userMatch)

	if result.Error != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), result.Error)
		return nil, customError.InternalServerError
	}

	return &userMatch, nil
}
func (m *UserModel) AcceptUserMatchToken(token string, home_id int, updated_at time.Time) (*UserMatchStruct, error) {
	result := database.DB.Table("tb_invitation").
		Where("tb_invitation.token = ? AND tb_invitation.home_id = ?", token, home_id).
		Updates(map[string]interface{}{
			"is_accept":  "A",
			"updated_at": updated_at,
		})

	if result.Error != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), result.Error)
		return nil, customError.InternalServerError
	}

	userMatch, err := m.GetUserMatch(token, home_id)
	if err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), result.Error)
		return nil, customError.InternalServerError
	}

	return userMatch, nil
}
