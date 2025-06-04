package main

import (
    "context"
    "fmt"
    "gorm_study/internal/config"
    "net/http"
)

func main() {
    if err := config.InitRedis("localhost:6379", "", 0); err != nil {
        panic(err)
    }
    defer config.CloseRedis()

    ctx := context.Background()
    if err := config.RedisClient.Set(ctx, "go-redis", "v9", 0).Err(); err != nil {
        panic(err)
    }
    var result string
    if err := config.RedisClient.Get(ctx, "go-redis").Scan(&result); err != nil {
        panic(err)
    }
    fmt.Println("redis test value:", result)

}

