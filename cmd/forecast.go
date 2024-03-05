/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

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

// forecastCmd represents the forecast command
var forecastCmd = &cobra.Command{
	Use:   "forecast",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

func init() {
	rootCmd.AddCommand(forecastCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// forecastCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// forecastCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func makeRequest() []byte {
	token := os.Getenv("CLIMA_TOKEN")
	if token == "" {
		fmt.Println("You need to declare an enviroment variable 'CLIMA_TOKEN' with your token. Create it here https://www.weatherapi.com/")
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
		fmt.Print(readErr)
	}
	return body
}
func printForecast(forecast Forecast) {
	fmt.Printf("%s, %s, %s\n", forecast.Location.Name, forecast.Location.Region, forecast.Location.Country)

	fmt.Printf("Temperature: %.1f°C\n", forecast.Current.TempC)
	fmt.Printf("Precipitation: %.1fmm\n", forecast.Current.PrecipMM)
	fmt.Printf("Condition: %s\n", forecast.Current.Condition.Text)
	fmt.Printf("Fells like: %.1f°C\n", forecast.Current.FellsLike)
	fmt.Printf("Humidity: %d%%\n", forecast.Current.Humidity)

	for _, day := range forecast.Forecast.Forecastday {
		fmt.Printf("Day: %s\n", day.Date)
		fmt.Printf("Condition: %s\n", day.Day.Condition.Text)
		for _, hour := range day.Hour {
			fmt.Printf("Time: %s, %.1f°C, %s, %.1f%%, %d\n", getTime(hour.Time), hour.TempC, hour.Condition.Text, hour.ChanceOfRain, hour.WillItRain)
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
