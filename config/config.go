package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
}

func LoadConfig() Config {
    v := viper.New()

    v.SetConfigName(".env")
    v.SetConfigType("env")
    v.AddConfigPath(".")
    v.AddConfigPath("..") // handle running from cmd/

    v.AutomaticEnv()

    // Defaults (in case env/.env missing)
    v.SetDefault("DB_HOST", "127.0.0.1")
    v.SetDefault("DB_PORT", "5432")
    v.SetDefault("DB_USER", "postgres")
    v.SetDefault("DB_PASSWORD", "root")
    v.SetDefault("DB_NAME", "GoNotebook_DB")
    v.SetDefault("JWT_SECRET", "supersecretkey")

    if err := v.ReadInConfig(); err != nil {
        log.Println("⚠️ No .env file found, relying only on environment variables")
    }

    return Config{
        DBHost:     v.GetString("DB_HOST"),
        DBPort:     v.GetString("DB_PORT"),
        DBUser:     v.GetString("DB_USER"),
        DBPassword: v.GetString("DB_PASSWORD"),
        DBName:     v.GetString("DB_NAME"),
        JWTSecret:  v.GetString("JWT_SECRET"),
    }
}
