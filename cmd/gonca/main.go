package main

//go:generate go run ../gqlgen/main.go

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/erdaltsksn/gonca/generated"
	"github.com/erdaltsksn/gonca/graph"
	"github.com/erdaltsksn/gonca/model"
)

func main() {
	// Connect the database
	dsn := "host=postgres user=gonca_user password=gonca_password dbname=gonca_db port=5432 sslmode=disable TimeZone=Asia/Istanbul"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.User{})

	// Connect to Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0, // use default DB
	})

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(
		generated.Config{
			Resolvers: &graph.Resolver{
				DB:    db,
				Redis: rdb,
			},
		},
	))

	http.Handle("/", srv)

	log.Fatal(http.ListenAndServe(":4000", nil))
}
