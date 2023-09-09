package controllers

import (
	"encoding/json"
	"itmsapp/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllData(c *gin.Context) {
	var hdata models.Hourly
	resp, err := http.Get("https://api.open-meteo.com/v1/forecast?latitude=11.7117117&longitude=79.3271609&timezone=IST&hourly=temperature_2m&hourly=relativehumidity_2m&hourly=windspeed_10m&hourly=winddirection_10m&hourly=pressure_msl&hourly=soil_temperature_6cm&hourly=visibility&hourly=rain")
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&hdata)
	if err != nil {
		log.Fatal(err)
	}
	models.DB.Find(&hdata)

	c.JSON(http.StatusOK, gin.H{"data": hdata})
}
