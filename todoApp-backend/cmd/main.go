package main

import (
	"fmt"
	"log"
	"os"
	"todoApp-backend/cmd/config"
	"todoApp-backend/src/Core/infrastructure/Web/controllers"
)

func main() {

	fmt.Println("Starting program")

	config.ReadEnv()

	fmt.Println("config loaded")

	fmt.Println("getting connection to DataBase")
	db := config.GetDB()
	fmt.Println("Connected to database")
	defer db.Close()

	templatesDir := os.Getenv("TEMPLATES_DIR")

	http := config.NewHttp(templatesDir)

	controller := controllers.NewController(db, http)

	controller.MountEndpoints()

	fmt.Println("Starting server")

	if err := http.Start(os.Getenv("APP_HOST") + ":" + os.Getenv("APP_PORT")); err != nil {
		log.Fatal(err)
	}
}
