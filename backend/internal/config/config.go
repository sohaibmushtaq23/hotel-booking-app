package config

// func GetConnectionString() string {
// 	return "sqlserver://sa:14845@DESKTOP-KD9VA3V/SQLEXPRESS?database=ClientDB&TrustServerCertificate=true"
// }

import (
	"fmt"
	"os"
)

func GetConnectionString() string {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	database := os.Getenv("DB_NAME")

	return fmt.Sprintf(
		"sqlserver://%s:%s@%s?database=%s&TrustServerCertificate=true",
		user,
		password,
		host,
		database,
	)
}
