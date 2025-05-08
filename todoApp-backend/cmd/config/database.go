package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"sync"
	"todoApp-backend/src/external/DataBase"
)

var (
	db   *sql.DB
	once sync.Once
	err  error
)

func connectDB(user, password, dbName, host, port, sslmode string) {
	fmt.Println("starting once.do")
	once.Do(func() {
		fmt.Println("creating DSN")
		connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, password, host, port, dbName, sslmode)
		fmt.Println("Starting sql.Open")
		db, err = sql.Open("postgres", connStr)
		fmt.Println("Database opened")
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		fmt.Println("pinning DB")
		err = db.Ping()
		if err != nil {
			log.Fatalf("Failed to ping database: %v", err)
		}
		fmt.Println("database pinned")

		log.Println("Successfully connected to database")

		migrator := DataBase.NewMigrator(db)

		if err := migrator.Migrate(); err != nil {
			log.Fatalf("error while migrating database: %v", err)
		}

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
