package main

import (
	"backend/internal/api/routes"
	"backend/internal/db"
)

func main() {
	db.GetDB()
	routes.GetRouter()
}
