package models

import (
	"time"

	"gorm.io/gorm"
)

type Hourly struct {
	gorm.Model
	Wis_Id              uint    `json:"wis_id" gorm:"primary_key"`
	Device_num          int     `json:"device_num"`
	Device_Name         string  `json:"device_name"`
	Date_Time           string  `json:"date&time" binding:"required"`
	Wind_Speed          int     `json:"wind_speed"`
	Wind_Speed_Kmh      float64 `json:"wind_speed_kmh"`
	Wind_Direction      float64 `json:"wind_direction"`
	Wind_cardinal       float64 `json:"wind_cardinal"`
	Air_Temperature     float64 `json:"air_temperature"`
	WeatherStation      string  `json:"weatherstation"`
	Relative_humidity   float64 `json:"relative_humidity"`
	Atmosphere_Pressure float64 `json:"atmosphere_pressure"`
	Visibility          float64 `json:"visibility"`
	Irst_Pav            float64 `json:"irst_pav"`
	Battery_Values      float64 `json:"battery_values"`
	Rain                float64 `json:"rain"`
	Road_Temperature    float64 `json:"road_temperature"`
	Date_added          string  `json:"date_added"`
}

type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
