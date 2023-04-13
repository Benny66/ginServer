// config/config.go

package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Config *config

type config struct {
	Mode      string
	IPAddress string
	Port      string

	DBHost     string
	DBPort     string
	DBUsername string
	DBPassword string
	DBName     string

	//redis
	RedisAddr     string
	RedisPassword string

	LogExpire   int //日志保存时间(单位：天)
	AppName     string
	AppVersion  string
	TokenSecret string
	Language    string
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err.Error())
	}

	Config = &config{
		Mode:      getEnv("MODE", "debug"),
		IPAddress: getEnv("IP_ADDRESS", "0.0.0.0"),
		Port:      getEnv("PORT", "8080"),

		DBHost:     getEnv("DB_HOST", ""),
		DBPort:     getEnv("DB_PORT", ""),
		DBUsername: getEnv("DB_USERNAME", ""),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", ""),

		RedisAddr:     getEnv("DB_REDIS_ADDRESS", ""),
		RedisPassword: getEnv("DB_REDIS_PASSWORD", ""),
		AppName:       getEnv("APP_NAME", "gin-server"),
		AppVersion:    getEnv("APP_VERSION", "v1.0.0"),
		TokenSecret:   getEnv("TOKEN_SECRET", "akhdfijwfwsefsdf"),
		Language:      getEnv("LANGUAGE", "zh-cn"),
	}
	Config.LogExpire, _ = strconv.Atoi(getEnv("LOG_EXPIRE", "30")) //日志保存时间(单位：天")

}

func getEnv(key string, def string) string {
	value := os.Getenv(key)
	if value == "" {
		return def
	}
	return value
}
