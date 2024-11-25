package main

import (
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/src/v1/common"
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/utils"
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/utils/database"
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/utils/log"
	customLog "github.com/GooDu-Dev/acuitmesh-intern-quiz/utils/log"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load("../../.env")
	customLog.InitLogger()
	database.ConDB()
}

func main() {
	err := database.DB.AutoMigrate(&database.UserModel{})
	if err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
	}

	database.DB.AutoMigrate(&database.HistoryModel{})
	if err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
	}
	database.DB.AutoMigrate(&database.InvitationModel{})
	if err != nil {
		log.Logging(utils.EXCEPTION_LOG, common.GetFunctionWithPackageName(), err)
	}
}
