package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "time"
)

type PostData struct {
    cmd   string `form:"cmd" json:"cmd" xml:"cmd" binding:"required"`
    key   string `form:"key" json:"key" xml:"key" binding:"required"`
    value string `form:"value" json:"value" xml:"value" binding:"required"`
}

func main() {
    app := gin.Default()

    app.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    app.POST("/testRedis.php", testRedis)

    app.Run()
}

func testRedis(c *gin.Context) {
    var p PostData
    var r RedisConn
    var re string

    if err := c.ShouldBind(&p); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    r.Init()
    defer r.Conn.Close()

    cmd   := c.PostForm("cmd")
    key   := c.PostForm("key")
    value := c.PostForm("value")

    switch cmd {
    case "set":
        SetRedisData(r.Conn, key, value, time.Duration(0))
    case "get":
        re = GetRedisData(r.Conn, key)
    case "del":
        DelRedisData(r.Conn, key)
    default:
        re = "happy"
    }

    c.JSON(http.StatusOK, gin.H{
        "errorCode": "0",
        "key":       key,
        "value":     value,
        "success":   re,
    })
}