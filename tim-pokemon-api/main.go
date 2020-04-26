package main

import (
	"fmt"
	"os"
	"yuting/daily-pg/2020-04-25/rest"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/joho/godotenv/autoload"
)

var (
	dbUrl  string
	dbName string
)

func init() {
	host := os.Getenv("HOST")
	fmt.Print(host)
	user := os.Getenv("USER")
	password := os.Getenv("PSSWORD")
	pgDbName := os.Getenv("PG_DB_NAME")
	dbUrl = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable ",
		host, user, password, pgDbName)
	fmt.Print(dbUrl)
	dbName = "postgres"
}

// @title Yuting-One-Month Project API
// @version 1.0
// @description Pokemon For Life  api server

// @contact.name yuting(Tim)
// @contact.email tim.wu@dotzero.tech

// @BasePath /v1
func main() {
	rest.RunAPI(dbName, dbUrl)
}
