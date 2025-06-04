package config

import (
    "context"
    "fmt"
    "github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis(host, password string, db int) error {
    RedisClient = redis.NewClient(&redis.Options{
        Addr:     host,
        Password: password,
        DB:       db,
    })

    if err := RedisClient.Ping(context.Background()).Err(); err != nil {
        return fmt.Errorf("Redis连接失败: %v", err)
    }
    fmt.Println("ping successful")
    return nil
}

func CloseRedis() {
    if RedisClient != nil {
        _ = RedisClient.Close()
    }
}

