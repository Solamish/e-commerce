package main

import (
	"e-commerce/models"
	"e-commerce/router"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	db, err := models.InitDB()
	if err != nil {
		log.Println("err open databases", err)
		return
	}
	defer db.Close()

	gin := gin.New()

	router.LOAD(
		gin,
	)

	gin.Run()
}
