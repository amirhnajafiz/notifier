package cache

import "github.com/go-redis/redis/v8"

func New(cfg Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.Host,     // "localhost:6379"
		Password: cfg.Password, // no password set
		DB:       0,            // use default DB
	})
}
