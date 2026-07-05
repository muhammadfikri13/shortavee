package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	// 1. Ambil DATABASE_URL dari environment variable (untuk di Render)
	dsn := os.Getenv("DB_HOST")

	// 2. Jika DATABASE_URL kosong (berarti Anda sedang jalankan di komputer lokal)
	//    Maka gunakan DSN localhost lama Anda sebagai cadangan (fallback)
	if dsn == "" {
		dsn = "host=localhost user=shortavee password=tinyavee123 dbname=shortavee_db port=5432 sslmode=disable"
	}

	// 3. GORM otomatis mengenali apakah dsn berupa format URL postgres:// atau format host=localhost
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
