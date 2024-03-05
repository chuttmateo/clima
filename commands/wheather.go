package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

type Forecast struct {
	Location struct {
		Name    string `json:"name"`
		Region  string `json:"region"`
		Country string `json:"country"`
	}
	Current struct {
		TempC     float64 `json:"temp_c"`
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
				Time  string  `json:"time"`
				TempC float64 `json:"temp_c"`
			}
		}
	}
}

var CmdWeather = &cobra.Command{
	Use: "forecast",
	Run: func(cmd *cobra.Command, args []string) {

		body := makeRequest()
		//body := getBodyFromFile()

		var forecast Forecast

		err := json.Unmarshal(body, &forecast)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			return
		}
		printForecast(forecast)

	},
}

func makeRequest() []byte {
	token := os.Getenv("CLIMA_TOKEN")
	if token == "" {
		fmt.Println("You need to declare an enviroment variable 'CLIMA_TOKEN' with your token. Here https://www.weatherapi.com/")
	}

	// URL for the API call
	url := "http://api.weatherapi.com/v1/forecast.json?key=" + token + "&q=Maria%20Grande,%20Entre%20Rios,%20Argentina&days=3&aqi=no&alerts=no"

	// Make the GET request
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return nil
	}
	defer response.Body.Close()

	body, readErr := io.ReadAll(response.Body)
	if readErr != nil {
		fmt.Print(err.Error())
	}
	return body
}
func printForecast(forecast Forecast) {
	fmt.Printf("%s, %s, %s\n", forecast.Location.Name, forecast.Location.Region, forecast.Location.Country)

	fmt.Printf("Temperature: %.1f°C\n", forecast.Current.TempC)
	fmt.Printf("Precipitation: %.1fmm\n", forecast.Current.PrecipMM)
	fmt.Printf("Condition: %s\n", forecast.Current.Condition.Text)

	for _, day := range forecast.Forecast.Forecastday {
		fmt.Printf("Day: %s\n", day.Date)
		fmt.Printf("Condition: %s\n", day.Day.Condition.Text)
		for _, hour := range day.Hour {
			fmt.Printf("Time: %s, %.1f°C, %s\n", getTime(hour.Time), hour.TempC, hour.Condition.Text)
		}
	}
}
func getTime(time string) string {
	s := strings.Split(time, " ")
	return s[1]
}
func getBodyFromFile() []byte {
	file, err := os.Open("commands/x.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return nil
	}
	defer file.Close()

	byteFile, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return nil
	}

	return byteFile
}
