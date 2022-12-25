package main

import (
	"fmt"
	"log"
	"wxcloudrun-golang/internal/app/service"
	"wxcloudrun-golang/internal/pkg/db"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}
	service := service.NewService()
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/auth/login", service.WeChatLogin)
	router.GET("/court", service.GetCounts)
	router.GET("/court/:id", service.GetCountInfo)

	router.POST("/event", service.StartEvent)
	router.POST("/collect/:fileID", service.ToggleCollectVideo)
	router.GET("/collect/user", service.GetCollectVideos)

	router.GET("/event/user", service.GetEventVideos)

	log.Fatal(router.Run())
}
