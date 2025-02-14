package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"sync"
)

var (
	db   *sql.DB
	once sync.Once
	err  error
)

func connectDB(user, password, dbName, host, port, sslmode string) {
	once.Do(func() {
		connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, password, host, port, dbName, sslmode)
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("Failed to ping database: %v", err)
		}

		log.Println("Successfully connected to database")
	})
}

func GetDB() *sql.DB {

	user := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSL_MODE")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")

	connectDB(user, password, dbName, host, port, sslmode)

	return db
}
