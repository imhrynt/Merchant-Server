package database

import (
	"fmt"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

type Config struct {
	PUBLIC_HOST   string
	PUBLIC_PORT   string
	DB_CONNECTION string
	DB_USERNAME   string
	DB_PASSWORD   string
	DB_PROTOCOL   string
	DB_ADDRESS    string
	DB_NAME       string
}

func NewENV() Config {
	return Config{
		PUBLIC_HOST:   getEnv("PUBLIC_HOST", "http://192.168.1.205"),
		PUBLIC_PORT:   getEnv("PUBLIC_PORT", "8080"),
		DB_CONNECTION: getEnv("DB_CONNECTION", "mysql"),
		DB_USERNAME:   getEnv("DB_USERNAME", "root"),
		DB_PASSWORD:   getEnv("DB_PASSWORD", ""),
		DB_PROTOCOL:   getEnv("DB_PROTOCOL", "tcp"),
		DB_ADDRESS:    getEnv("DB_ADDRESS", fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306"))),
		DB_NAME:       getEnv("DB_NAME", "merchant_db"),
	}
}

func NewDSN(env Config) *mysql.Config {
	config := mysql.NewConfig()
	return &mysql.Config{
		User:                 env.DB_USERNAME,
		Passwd:               env.DB_PASSWORD,
		Net:                  env.DB_PROTOCOL,
		Addr:                 env.DB_ADDRESS,
		DBName:               env.DB_NAME,
		Loc:                  config.Loc,
		MaxAllowedPacket:     config.MaxAllowedPacket,
		Timeout:              5 * time.Second,
		ReadTimeout:          1 * time.Second,
		WriteTimeout:         1 * time.Second,
		Logger:               config.Logger,
		AllowNativePasswords: config.AllowNativePasswords,
		CheckConnLiveness:    config.CheckConnLiveness,
		ParseTime:            true,
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
