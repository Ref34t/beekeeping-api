package main

type SenseBox struct {
	ID   string  `json:"id"`
	Temp float64 `json:"temperature"`
	Hum  float64 `json:"humidity"`
}
