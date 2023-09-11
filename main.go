package main

import (
	"itmsapp/controllers"
	"itmsapp/models"

	"github.com/gin-gonic/gin"
	// new
)

func main() {
	r := gin.New()
	//routes.UserRoute(r)
	r.Run("localhost:8081")
	models.ConnectDatabase() // new

	r.GET("/api/weather_data", controllers.GetAllData)
	r.POST("/api/weather_data", controllers.CreateData)

	err := r.Run()
	if err != nil {
		return
	}
}
