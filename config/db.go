package config

import (
	"fmt"
	"gorm.io/gorm"
)

var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "postgres",
		DBName:   "postgres",
	}
	return &dbConfig
}

func DbDSN(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Seoul",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DBName,
		dbConfig.Port,
	)
}