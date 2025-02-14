package main

import (
	"log"
	"os"
	"todoApp-backend/cmd/config"
)

func main() {

	config.ReadEnv()

	db := config.GetDB()
	defer db.Close()

	http := config.NewHttp()

	if err := http.Start(os.Getenv("APP_HOST") + ":" + os.Getenv("APP_PORT")); err != nil {
		log.Fatal(err)
	}
}
