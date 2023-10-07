package main

import (
	"backend/db"
	"fmt"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/joho/godotenv"
)

func main() {
	env := ".env.dev"
	
	err := godotenv.Load(env)
	if err != nil {
		panic("Error loading .env file")
	}
	err = rollbackMigrations()
	if err != nil {
		panic(err)
	}
	db.GetDB().Close()

	fmt.Println("Rollback completed successfully")
}

func rollbackMigrations() error {
	relativePath, err := filepath.Rel(".", "db/migrations")
	if err != nil {
		panic(err)
	}

	driver, err := postgres.WithInstance(db.GetDB(), &postgres.Config{})

	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+relativePath,
		"postgres", // Use the database driver
		driver,
	)

	if err != nil {
		return err
	}

	// Rollback all applied migrations
	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
