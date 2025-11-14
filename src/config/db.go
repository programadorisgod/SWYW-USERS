package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {

	dbUser := Env.DBUser
	password := Env.DBAppUserPassword
	host := Env.DBHostAuth
	dbname := Env.DBNameAuth
	port := Env.DBPort

	if dbUser == "" || password == "" || host == "" || dbname == "" || port == "" {
		fmt.Print(dbUser, password, dbUser, dbname, port)
		log.Fatal("❌ Some variables not found.")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, dbUser, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		fmt.Println("❌ Could not connect to DB: %v", err)
	}

	err = db.Ping()

	if err != nil {
		fmt.Println("❌ Could not ping to DB: %v", err)
	}

	fmt.Println("Successfully connected!")

	DB = db
}
