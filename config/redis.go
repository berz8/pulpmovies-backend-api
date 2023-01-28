package config

import (
	"context"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

var Rdb *redis.Client

func RedisClient() {
  db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
  if err != nil {
    panic("Redi db not valid")
  }

 Rdb = redis.NewClient(&redis.Options{
    Addr:     os.Getenv("REDIS_HOST"),
    Password: os.Getenv("REDIS_PASSWORD"), 
    DB:       db,  // use default DB
  })

}
