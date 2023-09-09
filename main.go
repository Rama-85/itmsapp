package main

import (
	"itmsapp/controllers"
	"itmsapp/models"

	"github.com/gin-gonic/gin"
	// new
)

func main() {
	r := gin.Default()

	models.ConnectDatabase() // new

	r.GET("/hdata", controllers.GetAllData)

	r.Run("localhost:8081")
	err := r.Run()
	if err != nil {
		return
	}
}
