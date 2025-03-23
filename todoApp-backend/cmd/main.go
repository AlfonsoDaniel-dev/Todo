package main

import (
	"fmt"
	"log"
	"os"
	"todoApp-backend/cmd/config"
)

func main() {

	fmt.Println("Starting program")

	config.ReadEnv()

	fmt.Println("config loaded")

	fmt.Println("getting connection to DataBase")
	db := config.GetDB()
	fmt.Println("Connected to database")
	defer db.Close()

	http := config.NewHttp()

	fmt.Println("Starting server")

	if err := http.Start(os.Getenv("APP_HOST") + ":" + os.Getenv("APP_PORT")); err != nil {
		log.Fatal(err)
	}
}
