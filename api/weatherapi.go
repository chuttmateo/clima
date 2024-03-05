package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Forecast struct {
	Location struct {
		Name    string `json:"name"`
		Region  string `json:"region"`
		Country string `json:"country"`
	}
	Current struct {
		TempC     float64 `json:"temp_c"`
		Humidity  int     `json:"humidity"`
		FellsLike float64 `json:"feelslike_c"`
		Condition struct {
			Text string `json:"text"`
		}
		PrecipMM float64 `json:"precip_mm"`
	}
	Forecast struct {
		Forecastday []struct {
			Date string `json:"date"`
			Day  struct {
				Condition struct {
					Text string `json:"text"`
				}
			}
			Hour []struct {
				Condition struct {
					Text string `json:"text"`
				}
				Time         string  `json:"time"`
				TempC        float64 `json:"temp_c"`
				ChanceOfRain float64 `json:"chance_of_rain"`
				WillItRain   int8    `json:"will_it_rain"`
			}
		}
	}
}

func makeRequest() ([]byte, error) {
	token := os.Getenv("CLIMA_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("You need to declare an environment variable 'CLIMA_TOKEN' with your token. Create it here https://www.weatherapi.com/")
	}

	// URL for the API call
	url := "http://api.weatherapi.com/v1/forecast.json?key=" + token + "&q=Maria%20Grande,%20Entre%20Rios,%20Argentina&days=3&aqi=no&alerts=no"

	// Make the GET request
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func GetForecast() (Forecast, error) {
	body, err := makeRequest()

	var forecast Forecast

	jsonErr := json.Unmarshal(body, &forecast)
	if jsonErr != nil {
		fmt.Println("Error parsing JSON:", jsonErr)
	}

	return forecast, err
}
