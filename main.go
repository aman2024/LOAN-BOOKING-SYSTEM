package main

import (
	"loan-booking/db"
	"loan-booking/routes"
	"loan-booking/services"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {

	LoadEnvVariables()

	dbClient, err := db.InitMySQL()
	if err != nil {
		panic(err)
	}
	services := services.Services{
		DB: &db.DbClient{Client: dbClient},
	}
	router := routes.Init(&services)
	if err := router.Run(":8000"); err != nil {
		panic(err)
	}
}

func LoadEnvVariables() {
	env := os.Getenv("ENV")
	if strings.ToUpper(env) == "DEV" {
		err := godotenv.Load(".env.test")
		if err != nil {
			log.Fatal("Error loading .env file with error: ", err)
		}
	}
}
