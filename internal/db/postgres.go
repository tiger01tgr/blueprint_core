package db

import (
	"os"
	"database/sql"
	"fmt"
	"sync"
	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

func GetDB() *sql.DB {
	once.Do(Init)
	return db
}

func Init() {
	db = connect()
}

func connect() *sql.DB {
	db_password := " password=" + os.Getenv("DB_PASSWORD")
	db_host := " host=" + os.Getenv("DB_HOST")
	db_port := " port=" + os.Getenv("DB_PORT")
	connection := "postgresql://postgres:" + db_password + "@" + db_host + ":" + db_port + "/postgres"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Printf("Error pinging to database: %v\n", err)
		panic(err)
	}
	return db
}
