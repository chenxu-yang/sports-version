package main

import (
	"fmt"
	"log"
	"net/http"
	"wxcloudrun-golang/app/service"
	"wxcloudrun-golang/pkg/db"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}

	router := gin.Default()
	router.GET("/auth/login", service.WeChatLogin)
	log.Fatal(http.ListenAndServe(":80", nil))
}
