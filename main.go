package main

import (
	"ryuichi.com/go-bank-backend/api"
	"ryuichi.com/go-bank-backend/database"
)

func main() {
	database.InitDatabase()
	api.StartApi()
}