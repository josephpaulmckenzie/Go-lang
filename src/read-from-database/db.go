package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var (
		host     = os.Getenv("host")
		port     = os.Getenv("port")
		user     = os.Getenv("user")
		password = os.Getenv("password")
		dbname   = "joesdatabase"
	)
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	rows, err := db.Query("SELECT username,given_name,fcmtoken,email_address FROM usertable")
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var username string
		var fcmtoken string
		var givenName string
		var emailAddress string
		err = rows.Scan(&username, &givenName, &fcmtoken, &emailAddress)
		if err != nil {
			// handle this error
			panic(err)
		}

		fmt.Println(username, givenName, fcmtoken, emailAddress)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	// fmt.Println("Successfully connected!")
}
