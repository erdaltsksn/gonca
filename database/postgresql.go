package database

import (
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PostgreSQL returns a PostgreSQL connection pool.
func PostgreSQL() *gorm.DB {
	var dsn string
	if viper.IsSet("postgresql.host") {
		dsn += "host=" + viper.GetString("postgresql.host") + " "
	}
	if viper.IsSet("postgresql.port") {
		dsn += "port=" + viper.GetString("postgresql.port") + " "
	}
	if viper.IsSet("postgresql.user") {
		dsn += "user=" + viper.GetString("postgresql.user") + " "
	}
	if viper.IsSet("postgresql.pass") {
		dsn += "password=" + viper.GetString("postgresql.pass") + " "
	}
	if viper.IsSet("postgresql.name") {
		dsn += "dbname=" + viper.GetString("postgresql.name") + " "
	}
	dsn += "sslmode=disable TimeZone=Asia/Istanbul"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
