package main

import (
	"skyshi-rest-api/config"
)

func main() {
	db, _ := config.DBConnection()
	config.Setup(db)
}
