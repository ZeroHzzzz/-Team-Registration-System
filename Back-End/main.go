package main

import (
	"backend/app/database"
	"backend/app/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()

	r := gin.Default()

	router.Init(r)
	err := r.Run()
	if err != nil {
		log.Fatal("Server start error", err)
	}
}
