package main

import (
	"github.com/GooDu-Dev/acuitmesh-intern-quiz/utils/database"
	customLog "github.com/GooDu-Dev/acuitmesh-intern-quiz/utils/log"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load("../../.env")
	customLog.InitLogger()
	database.ConDB()
}

func main() {
	// database.DB.AutoMigrate(&database.BuildingModel{})
}
