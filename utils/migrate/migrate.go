package main

import (
	"github.com/GooDu-dev/acuitmesh-intern-quiz/utils/database"
	customLog "github.com/GooDu-dev/acuitmesh-intern-quiz/utils/log"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load("../../.env")
	customLog.InitLogger()
	database.ConDB()
}

func main() {
	database.DB.AutoMigrate(&database.BuildingModel{})
	database.DB.AutoMigrate(&database.AreaModel{})
	database.DB.AutoMigrate(&database.AreaTypeModel{})
}
