package main

import (
	"backend/api/routes"
	"backend/config/firebase"
	"backend/config/s3Client"
	"backend/db"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	var env string
	if os.Getenv("ENV") == "prod" {
		env = ".env"
	} else if os.Getenv("ENV") == "dev" {
		env = ".env.dev"
	} else if os.Getenv("ENV") == "deploy" {
		
	}

	err := godotenv.Load(env)
	if err != nil {
		panic("Error loading .env file")
	}
	log.Println(env)
	db.GetDB()
	defer db.GetDB().Close()
	firebase.GetFirebase()
	s3Client.GetS3()
	
	routes.GetRouter()
}
