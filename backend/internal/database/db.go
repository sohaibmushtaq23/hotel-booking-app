package database

import (
	"database/sql"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var DB *sql.DB

func Connect(connString string) {
	var err error
	DB, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("DB Open Error:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("DB Ping Error:", err)
	}

	log.Println("Connected to SQL Server successfully")
}
