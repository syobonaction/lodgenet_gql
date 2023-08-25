package graph

import (
	"context"

	"netreality.world/m/graph/datasources"
	"netreality.world/m/graph/model"
)

// CurrentWeather is the resolver for the currentWeather field.
func (r *queryResolver) CurrentWeather(ctx context.Context, lat string, lon string) (*model.CurrentWeather, error) {
	d := &datasources.CurrentWeather{}
	err := datasources.GetCurrentWeather(d, lat, lon)

	if err != nil {
		return nil, err
	} else {
		w := []*model.Weather{}

		for _, s := range d.Weather {
			e := &model.Weather{
				ID:          s.Id,
				Type:        s.Main,
				Description: s.Description,
			}

			w = append(w, e)
		}

		a := &model.Atmosphere{
			Temperature: &model.Temperature{
				Real:      d.Main.Temp,
				Min:       d.Main.Temp_min,
				Max:       d.Main.Temp_max,
				Feelslike: d.Main.Feels_like,
			},
			Pressure: &d.Main.Pressure,
			Humidity: &d.Main.Humidity,
		}

		c := &model.Conditions{
			Wind: &model.Wind{
				Speed:  d.Wind.Speed,
				Degree: d.Wind.Deg,
				Gust:   d.Wind.Gust,
			},
			Sunrise: d.Sys.Sunrise,
			Sunset:  d.Sys.Sunset,
		}

		m := &model.CurrentWeather{
			Weather:    w,
			Atmosphere: a,
			Conditions: c,
		}

		return m, nil
	}
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
