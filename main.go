package main

import (
	"fmt"
	"log"
	"net/http"
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

	router.GET("/auth/login", service.WeChatLogin)
	log.Fatal(http.ListenAndServe(":80", nil))
}
