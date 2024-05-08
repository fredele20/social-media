package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)


type Secrets struct {
	DatabaseURL string `json:"DATABASE_URL"`
	DatabaseName string `json:"DATABASE_NAME"`
	JwtSecretKey string `json:"JWT_SECRETKEY"`
	HttpPort string
}

var s Secrets

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file")
	}

	s = Secrets{}

	s.DatabaseURL = os.Getenv("DATABASE_URL")
	s.DatabaseName = os.Getenv("DATABASE_NAME")
	s.JwtSecretKey = os.Getenv("JWT_SECRETKEY")

	if s.HttpPort = os.Getenv("HTTP_PORT"); s.HttpPort == "" {
		s.HttpPort = "80"
	}
}

func GetSecrets() Secrets {
	return s
}
