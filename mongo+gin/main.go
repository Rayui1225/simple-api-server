package main

import (
    "github.com/gin-gonic/gin"
    "context"
)

func main() {
    router := gin.Default()
    mgoCli := connectDB()
    defer mgoCli.Disconnect(context.TODO()) // 确保在退出main函数时断开连接

    AdvertiseHandler := NewAdvertiseHandler(mgoCli)
    router.GET("/Advertise", AdvertiseHandler.GetAdvertises)
    router.GET("/Ad", AdvertiseHandler.getAdvertiseByCondition)
    router.POST("/Advertise", AdvertiseHandler.postAdvertises)
    router.Run("localhost:8080")
}
