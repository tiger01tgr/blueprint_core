package main

import (
	"backend/api/routes"
	"backend/config"
	"backend/db"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	var env string
	if os.Getenv("ENV") == "prod" {
		env = ".env"
	} else {
		env = ".env.dev"
	}

	err := godotenv.Load(env)
	if err != nil {
		panic("Error loading .env file")
	}
	log.Println(env)
	db.GetDB()
	defer db.GetDB().Close()
	config.GetFirebase()
	
	
	routes.GetRouter()
}
