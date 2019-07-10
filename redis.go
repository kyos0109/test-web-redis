package main

import (
    "github.com/go-redis/redis"
    "time"
    "fmt"
)

type RedisConn struct {
    Conn *redis.Client
}

func (r *RedisConn) Init() {
    r.Conn = redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })

    pong, err := r.Conn.Ping().Result()
    fmt.Println(pong, err)
}

func GetRedisData(client *redis.Client, key string) string {
    val, err := client.Get(key).Result()
    if err != nil {
        defer client.Close()
        panic(err)
    }

    return val
}

func SetRedisData(client *redis.Client, key ,value string, ttl time.Duration) {
    err := client.Set(key, value, ttl).Err()
    if err != nil {
        defer client.Close()
        panic(err)
    }
}

func DelRedisData(client *redis.Client, key string) {
    err := client.Del(key).Err()
    if err != nil {
        defer client.Close()
        panic(err)
    }
}