package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	User     string
	Pass     string
	Host     string
	Port     string
	Name     string
	SslMode  string
	Timezone string
}

// NewDB は、DB接続の初期化を実行します
func NewDB() (*gorm.DB, error) {
	cfg := DBConfig{
		User: os.Getenv("DB_USER"),
		Pass: os.Getenv("DB_PASS"),
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		Name: os.Getenv("DB_NAME"),
		SslMode: os.Getenv("DB_SSLMODE"),
		Timezone: os.Getenv("DB_TIMEZONE"),
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", cfg.Host, cfg.User, cfg.Pass, cfg.Name, cfg.Port, cfg.SslMode, cfg.Timezone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, err
}
