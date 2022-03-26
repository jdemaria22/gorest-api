package config

import (
	"log"

	"github.com/joho/godotenv"
)

func InitConfig() {
	err := godotenv.Load("envoirment.env")
	if err != nil {
		log.Fatal(err)
	}
}
