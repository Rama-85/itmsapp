package controllers

import (
	"fmt"
	"itmsapp/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WeatherData struct {
	DeviceNum          int     `json:"device_num"  binding:"required"`
	DeviceName         string  `json:"device_name"  binding:"required"`
	DateTime           string  `json:"date_time" binding:"required"`
	WindSpeed          int     `json:"wind_speed"  binding:"required"`
	WindSpeed_Kmh      float64 `json:"wind_speed_kmh"  binding:"required"`
	WindDirection      float64 `json:"wind_direction"  binding:"required"`
	Windcardinal       float64 `json:"wind_cardinal"  binding:"required"`
	AirTemperature     float64 `json:"air_temperature"  binding:"required"`
	WeatherStation     string  `json:"weatherstation"  binding:"required"`
	Relativehumidity   float64 `json:"relative_humidity"  binding:"required"`
	AtmospherePressure float64 `json:"atmosphere_pressure"  binding:"required"`
	Visibility         float64 `json:"visibility"  binding:"required"`
	IrstPav            float64 `json:"irst_pav"  binding:"required"`
	BatteryValues      float64 `json:"battery_values"  binding:"required"`
	Rain               float64 `json:"rain"  binding:"required"`
	RoadTemperature    float64 `json:"road_temperature"  binding:"required"`
	Dateadded          string  `json:"date_added"  binding:"required"`
}

func GetAllData(c *gin.Context) {
	var weather_data []WeatherData

	/*
		resp, err := http.Get("https://api.open-meteo.com/v1/forecast?latitude=11.7117117&longitude=79.3271609&timezone=IST&hourly=temperature_2m&hourly=relativehumidity_2m&hourly=windspeed_10m&hourly=winddirection_10m&hourly=pressure_msl&hourly=soil_temperature_6cm&hourly=visibility&hourly=rain")
		if err != nil {
			panic(err.Error())
		}
		defer resp.Body.Close()

		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&weather_data)
		if err != nil {
			log.Fatal(err)
		}
	*/
	err := models.DB.Find(&weather_data).Error
	if err != nil {
		fmt.Println("err while fetch : ", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": weather_data})
}

func CreateData(c *gin.Context) {
	// Validate input
	var input WeatherData
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//weather_data := models.Hourly{Date_Time: input.Date_Time, WeatherStation: input.WeatherStation, Air_Temperature: input.Air_Temperature, Battery_Values: input.Battery_Values, Relative_humidity: input.Relative_humidity, Road_Temperature: input.Road_Temperature, Wind_Speed: input.Wind_Speed, Wind_Direction: input.Wind_Direction, Visibility: input.Visibility, Rain: input.Rain}
	//models.DB.Create(&weather_data)

	c.JSON(http.StatusOK, gin.H{"data": input})
}
