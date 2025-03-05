package main

// SenseBox represents the data for a senseBox sensor
type SenseBox struct {
	ID   string  `json:"id"`
	Temp float64 `json:"temperature"`
	Hum  float64 `json:"humidity"`
}
