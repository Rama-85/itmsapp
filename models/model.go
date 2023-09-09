package models

import (
	"time"

	"gorm.io/gorm"
)

/*
	type OpenMeteoResponse struct {
		Id                    int          `json:"id" gorm:"primary_key"`
		Latitude              float64      `json:"latitude"`
		Longitude             float64      `json:"longitude"`
		Generationtime_ms     float64      `json:"generationtime_ms"`
		Utc_offset_seconds    int          `json:"utc_offset_seconds"`
		Timezone              string       `json:"timezone"`
		Timezone_abbreviation string       `json:"timezone_abbrevation"`
		Elevation             float64      `json:"elevation"`
		Hourly_units          Hourly_units `json:"hourly_units"`
		Hourly                Hourly       `json:"hourly"`
	}

	type Hourly_units struct {
		//Time                 string `json:"time"`
		Temperature_2m       string `json:"temperature_3m"`
		Relativehumidity_2m  string `json:"relativehumidity_2m"`
		Windspeed_10m        string `json:"windspeed_10m"`
		Winddirection_10m    string `json:"winddirection_10m"`
		Pressure_msl         string `json:"pressure_msl"`
		Soil_temperature_6cm string `json:"soil_temperature_6cm"`
		Visibility           string `json:"visibility"`
		Rain                 string `json:"rain"`
	}
*/
type Hourly struct {
	Id                   int     `json:"id" gorm:"primary_key"`
	Time                 string  `json:"time"`
	Temperature_2m       float64 `json:"temperature_2m"`
	Relativehumidity_2m  int     `json:"relativehumidity_2m"`
	Windspeed_10m        float64 `json:"windspeed_10m"`
	Winddirection_10m    int     `json:"winddirection_10m"`
	Pressure_msl         float64 `json:"pressure_msl"`
	Soil_temperature_6cm float64 `json:"soil_temperature_6cm"`
	Visibility           float64 `json:"visibility"`
	Rain                 float64 `json:"rain"`
}
type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
