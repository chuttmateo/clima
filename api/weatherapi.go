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

// makeRequest makes a GET request to the weather API and returns the response body.
//
// No parameters.
// Returns a byte slice and an error.
func makeRequest() ([]byte, error) {
	return makeRequestWithLocation("")
}

func makeRequestWithLocation(location string) ([]byte, error) {
	//get token environment variable
	token := os.Getenv("CLIMA_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("you need to declare an environment variable 'CLIMA_TOKEN' with your token. Create it here https://www.weatherapi.com/")
	}

	//get location environment variable if location flag is empty
	if location == "" {
		location = os.Getenv("CLIMA_LOCATION")
		if location == "" {
			return nil, fmt.Errorf("you need to declare an environment variable 'CLIMA_LOCATION' with your location. For example CLIMA_LOCATION='Maria Grande, Entre Rios, Argentina'")
		}
	}

	// make GET request to API to get user by ID
	apiUrl := "http://api.weatherapi.com/v1/forecast.json"
	request, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return nil, err
	}

	// Set query parameters
	query := request.URL.Query()
	query.Add("key", token)
	query.Add("q", location)
	query.Add("days", "3")
	query.Add("aqi", "no")
	query.Add("alerts", "no")
	request.URL.RawQuery = query.Encode()

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
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
	return GetForecastWithLocation("")
}
func GetForecastWithLocation(location string) (Forecast, error) {
	body, err := makeRequestWithLocation(location)

	var forecast Forecast

	jsonErr := json.Unmarshal(body, &forecast)
	if jsonErr != nil {
		fmt.Println("Error parsing JSON:", jsonErr)
	}

	return forecast, err
}
