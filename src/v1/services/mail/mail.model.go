package mail

import (
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/src/v1/common"
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/utils"
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/utils/database"
	customError "github.com/GooDu-Dev/acuitmesh-intern-quiz/utils/error"
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/utils/log"
)

type MailModel struct {
}

func (m *MailModel) InitModel() *MailModel {
	return &MailModel{}
}

func (m *MailModel) GetMailUser(user_id int) (*MailUser, error) {
	var mail MailUser

	result := database.DB.Table("tb_user").
		Select(
			"tb_user.username",
			"tb_user.mail as email",
		).
		Where("tb_user.id = ?", user_id).
		Find(&mail)

	if result.Error != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), result.Error)
		return nil, customError.InternalServerError
	}

	return &mail, nil
}
