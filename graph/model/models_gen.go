// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Current struct {
	Temperature   int    `json:"temperature"`
	Feelslike     int    `json:"feelslike"`
	WindSpeed     int    `json:"windSpeed"`
	WindDirection string `json:"windDirection"`
	IsDay         bool   `json:"isDay"`
}

type Location struct {
	Name    string   `json:"name"`
	Region  string   `json:"region"`
	Current *Current `json:"current"`
}
