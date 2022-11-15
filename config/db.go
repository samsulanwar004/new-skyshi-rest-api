package config

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConnection() (*gorm.DB, error) {

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("Some error occured. Err: %s", err)
	// }

	HOST := os.Getenv("MYSQL_HOST")
	PORT := os.Getenv("MYSQL_PORT")
	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_PASSWORD")
	DBNAME := os.Getenv("MYSQL_DBNAME")

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&sql_mode=''", USER, PASS, HOST, PORT, DBNAME)
	return gorm.Open(mysql.Open(url), &gorm.Config{})
}
