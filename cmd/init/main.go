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
	firebase.GetFirebase()
	s3Client.GetS3()
	
	routes.GetRouter()
}
