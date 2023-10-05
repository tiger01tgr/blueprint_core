package main

import (
	"backend/api/routes"
	"backend/config"
	"backend/db"
	"fmt"
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
	fmt.Println(env)
	db.GetDB()
	defer db.GetDB().Close()
	routes.GetRouter()
	config.GetFirebase()

	fmt.Println("Server is running.")
}
