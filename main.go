package main

import (
	"GIN-GORM-REST/controllers"
	"GIN-GORM-REST/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()
	r.GET("/books", controllers.FindBooks)
	r.Run()
}
