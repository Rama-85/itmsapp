package main

import (
	"itmsapp/controllers"
	"itmsapp/models"

	"github.com/gin-gonic/gin"
	// new
)

func main() {

	models.ConnectDatabase() // new

	r := gin.Default()
	//routes.UserRoute(r)

	r.GET("/api/weather_data", controllers.GetAllData)
	// r.POST("/api/input", controllers.CreateData)

	r.Run("localhost:8081")
	// err := r.Run()
	// if err != nil {
	// 	return
	// }
}
