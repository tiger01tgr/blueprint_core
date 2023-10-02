package db

import (
	"database/sql"
	"os"
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
	db_password, p_error := os.LookupEnv("DB_PASSWORD")
	db_host, h_error := os.LookupEnv("DB_HOST")
	db_port, port_error := os.LookupEnv("DB_PORT")
	if !p_error {
		panic("DB_PASSWORD not found")
	}
	if !h_error {
		panic("DB_HOST not found")
	}
	if !port_error {
		panic("DB_PORT not found")
	}
	connection := "postgresql://postgres:" + db_password + "@" + db_host + ":" + db_port + "/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
