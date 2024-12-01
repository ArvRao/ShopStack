package database

import (
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

// RedisClient is the global Redis client
var RedisClient *redis.Client
var Ctx = context.Background()

// InitRedis initializes the Redis connection
func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"), // e.g., "localhost:6379"
		Password: "",                      // No password set
		DB:       0,                       // Default DB
	})

	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Println("Connected to Redis successfully!")
}
