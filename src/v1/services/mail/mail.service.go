package mail

import (
	"os"
	"strconv"

	"github.com/GooDu-Dev/acuitmesh-intern-quiz/src/v1/common"
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/utils"
	customError "github.com/GooDu-Dev/acuitmesh-intern-quiz/utils/error"
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/utils/log"
	gomail "gopkg.in/gomail.v2"
)

type MailService struct {
	model *MailModel
}

var s *MailService

func GetService() *MailService {
	if s != nil {
		return s
	}
	return s.InitService()
}

func (s *MailService) InitService() *MailService {
	model := MailModel{}
	return &MailService{
		model: model.InitModel(),
	}
}

func (s *MailService) SendInvite(sender_id int, receiver_id int, token string) (err error) {
	var sender *MailUser
	var receiver *MailUser

	if sender, err = s.model.GetMailUser(sender_id); err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		return err
	}
	if receiver, err = s.model.GetMailUser(receiver_id); err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		return err
	}

	m := gomail.NewMessage()

	m.SetHeader("From", os.Getenv("PUBLIC_EMAIL_USERNAME"))
	m.SetHeader("To", receiver.Email)
	m.SetHeader("Subject", UserInvite.ThHeader(sender.Username))
	m.SetHeader("text/plain", UserInvite.ThMessage(sender.Username, receiver.Username, token))

	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		return customError.InternalServerError
	}

	d := gomail.NewDialer(os.Getenv("SMTP_DOMAIN"), port, os.Getenv("PUBLIC_EMAIL_USERNAME"), os.Getenv("PUBLIC_EMAIL_PWD"))

	if err := d.DialAndSend(m); err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
		return customError.InternalServerError
	}

	log.Logging(utils.INFO_LOG, common.GetFunctionWithPackageName(), "Email sent Successfully.")
	return nil

}
