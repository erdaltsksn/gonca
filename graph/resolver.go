package graph

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.

// Resolver serves as dependency injection, add any dependenc here.
type Resolver struct {
	DB    *gorm.DB
	Redis *redis.Client
}
