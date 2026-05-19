package config

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"os"
	"sigap-sultan-be/src/common"
	"strconv"
)

func ReadEnv() *common.EnvConfig {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf(": %s", err)
	}

	return &common.EnvConfig{
		AddrListen: os.Getenv("ADDR_LISTEN"),
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     int16(port),
		DbUsername: os.Getenv("DB_USERNAME"),
		DbPass:     os.Getenv("DB_PASSWORD"),
		DbName:     os.Getenv("DB_NAME"),
		DbScheme:   os.Getenv("DB_SCHEME"),
	}
}
