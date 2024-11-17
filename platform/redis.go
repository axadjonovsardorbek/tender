package platform

import (
	"context"
	"fmt"
	"log"

	"github.com/axadjonovsardorbek/tender/config"
	"github.com/redis/go-redis/v9"
)

// Redis wraps the Redis client
type Redis struct {
	Client *redis.Client
}

// ConnectRedis initializes the Redis connection
func ConnectRedis(cfg *config.Config) *Redis {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
	})

	// Use a context for the Redis operations
	ctx := context.Background()

	// Verify connection
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Connected to Redis.")
	return &Redis{Client: client}
}

// Close closes the Redis connection
func (r *Redis) Close() {
	if err := r.Client.Close(); err != nil {
		log.Printf("Error closing Redis: %v", err)
	}
}
