package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/khofesh/simple-go-api/common"
	"github.com/khofesh/simple-go-api/setup"
)

func main() {
	// load env files
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env files")
	}
	var mongoURI string = os.Getenv("MONGO_URI")

	if err = common.InitMongo(mongoURI); err != nil {
		log.Fatal(err.Error())
	}

	r := setup.Router()

	r.Run(":8090")
}
