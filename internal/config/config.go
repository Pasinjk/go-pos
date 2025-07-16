package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() error {
	currentDir, _ := os.Getwd()
	err := godotenv.Load(currentDir + "/.env")
	if err != nil {
		return err
	}

	return nil
}

func GetConfig() (string, error) {

	if err := loadEnv(); err != nil {
		return "", err
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	ssl := os.Getenv("DB_SSLMODE")

	if host == "" || port == "" || user == "" || password == "" || name == "" || ssl == "" {
		return "", errors.New("missing required database environment variables")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, name, ssl)

	return psqlInfo, nil
}
