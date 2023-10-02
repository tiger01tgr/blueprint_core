package main

import (
	"github.com/joho/godotenv"
	"backend/internal/api/routes"
	"backend/internal/db"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	db.GetDB()
	routes.GetRouter()
}
