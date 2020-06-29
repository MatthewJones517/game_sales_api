package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// DbConn is the variable the contains the database connection.
var DbConn *sql.DB

// SetupDatabase prepares the database connection for use.
func SetupDatabase() {
	// Pull in connection data from env variables
	err := godotenv.Load("config.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPass := os.Getenv("MYSQL_PASS")
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlDbnm := os.Getenv("MYSQL_DBNM")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s", mysqlUser, mysqlPass, mysqlHost, mysqlDbnm)

	DbConn, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	DbConn.SetMaxOpenConns(3)
	DbConn.SetMaxIdleConns(3)
	DbConn.SetConnMaxLifetime(60 * time.Second)
}
