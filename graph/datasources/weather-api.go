package datasources

import (
	"encoding/json"
	"net/http"
	"os"
)

var client http.Client

type Coordinates struct {
	Lat float64
	Lon float64
}

type Weather struct {
	Id          int
	Main        string
	Description string
	Icon        string
}

type Main struct {
	Temp       float64
	Feels_like float64
	Temp_min   float64
	Temp_max   float64
	Pressure   int
	Humidity   int
}

type Wind struct {
	Speed float64
	Deg   int
	Gust  float64
}

type Clouds struct {
	All int
}

type Sys struct {
	Type    int
	Id      int
	Country string
	Sunrise int
	Sunset  int
}

type CurrentWeather struct {
	Coordinates Coordinates
	Weather     []*Weather
	Base        string
	Main        Main
	Visibility  int
	Wind        Wind
	Clouds      Clouds
	Dt          int
	Sys         Sys
	Timezone    int
	Id          int
	Name        string
	Cod         interface{}
}

func GetCurrentWeather(target interface{}, lat string, lon string) error {
	api_key := os.Getenv("API_KEY")
	url := "https://api.openweathermap.org/data/2.5/weather?lat=" + lat + "&lon=" + lon + "&appid=" + api_key + ""
	resp, err := client.Get(url)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	t := json.NewDecoder(resp.Body).Decode(target)

	return t

}
