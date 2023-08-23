package datasources

import (
	"encoding/json"
	"net/http"
	"os"
)

var client http.Client

type Request struct {
	Type     string
	Query    string
	Language string
	Unit     string
}

type Location struct {
	Name        string
	Country     string
	Region      string
	Lat         string
	Lon         string
	Timezone_id string
}

type Current struct {
	Observation_time     string
	Temperature          int
	Weather_code         int
	Weather_icons        []string
	Weather_descriptions []string
	Wind_speed           int
	Wind_degree          int
	Wind_dir             string
	Pressure             int
	Precip               int
	Humidity             int
	Cloudcover           int
	Feelslike            int
	Uv_index             int
	Visibility           int
	Is_day               string
}

type WeatherData struct {
	Request  Request
	Location Location
	Current  Current
}

func GetWeather(target interface{}, zip string) error {
	api_key := os.Getenv("API_KEY")
	url := "http://api.weatherstack.com/current?access_key=" + api_key + "&query=" + zip + ""
	resp, err := client.Get(url)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	t := json.NewDecoder(resp.Body).Decode(target)

	return t
}
