package connection

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func LoadEnv(key string) string {
	return os.Getenv(key)
}

func GetConnection() (*gorm.DB, error) {

	dbHost := LoadEnv("DB_HOST")
	dbPort := LoadEnv("DB_PORT")
	dbUsername := LoadEnv("DB_USERNAME")
	dbPassword := LoadEnv("DB_PASSWORD")
	dbName := LoadEnv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	return db, nil
}
