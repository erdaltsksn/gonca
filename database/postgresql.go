package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PostgreSQL returns a PostgreSQL connection pool.
func PostgreSQL() *gorm.DB {
	dsn := "host=postgres user=gonca_user password=gonca_password dbname=gonca_db port=5432 sslmode=disable TimeZone=Asia/Istanbul"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
