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

	var user []UserCard

	result := database.DB.Table("tb_user").
		Select(
			"tb_user.id",
			"tb_user.username",
		).
		Where("tb_user.mail = ? AND tb_user.pwd = ?", email, pwd).
		Find(&user)

	if result.Error != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), result.Error)
		return nil, customError.InternalServerError
	}
	if len(user) == 0 {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), user)
		return nil, customError.UserNotFound
	}
	if len(user) > 1 {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), user)
		return nil, customError.InternalServerError
	}

	return &user[0], nil
}

func (m *UserModel) CreateUser(email string, username string, pwd string) (*UserCard, error) {

	user := UserCreateStruct{
		Username:  username,
		Email:     email,
		Pwd:       pwd,
		CreatedAt: time.Now(),
	}

	result := database.DB.Table("tb_user").
		Select("username", "mail", "pwd", "created_at").
		Create(&user)

	if result.Error != nil {
		if common.IsPostgresqlDataDup(result.Error) {
			log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), "data dupplicated")
			return nil, customError.UserAccountDupplicated
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
