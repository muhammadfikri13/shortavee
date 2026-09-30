package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	var dsn string

	// Prioritas 1: DATABASE_URL (Neon / Render)
	dsn = os.Getenv("DATABASE_URL")

	// Prioritas 2: PostgreSQL Kubernetes
	if dsn == "" {
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		dbname := os.Getenv("DB_NAME")

		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			host,
			user,
			password,
			dbname,
			port,
		)
	}

	// Fallback terakhir untuk local development
	if dsn == "" || dsn == "host= user= password= dbname= port= sslmode=disable" {
		dsn = "host=localhost user=shortavee password=Tinyavee123!@# dbname=shortavee_db port=5432 sslmode=disable"
	}

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
