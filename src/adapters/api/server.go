package api

import (
	"fmt"
	"go-viacep-wrapper/adapters/api/configs"
	"log"
	"os"
)

func Run() {
	app := configs.App()

	fmt.Printf("API is running at PORT %v", os.Getenv("PORT"))
	log.Fatal(app.Listen(":8080"))
}
