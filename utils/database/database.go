package database

import (
	"fmt"
	"os"

	"github.com/GooDu-Dev/acuitmesh-intern-quiz/src/v1/common"
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/utils"
	customLog "github.com/GooDu-Dev/acuitmesh-intern-quiz/utils/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConDB() {
	url := os.Getenv("DB_URL")
	fmt.Println("url", url)
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		customLog.Logging(utils.ERR_LOG, common.GetFunctionWithPackageName(), err)
		return
	} else {
		DB = db
		customLog.Logging(utils.INFO_LOG, common.GetFunctionWithPackageName(), "Connect to database successfully")
	}
}

type Tabler interface {
	TableName() string
}
