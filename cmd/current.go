/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/chuttmateo/clima/api"
	"github.com/spf13/cobra"
)

// currentCmd represents the current command
var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "The current wheather conditions",
	Long:  `This command will return the current wheather conditions.`,
	Run: func(cmd *cobra.Command, args []string) {

		forecast, err := api.GetForecast()
		if err != nil {
			fmt.Println("Error getting forecast:", err)
			return
		}
		printCurrent(forecast)

	},
}

func init() {
	rootCmd.AddCommand(currentCmd)
}

func printCurrent(forecast api.Forecast) {
	fmt.Printf("Current weather in: %s, %s, %s\n", forecast.Location.Name, forecast.Location.Region, forecast.Location.Country)

	fmt.Printf("Temperature: %.1f°C\n", forecast.Current.TempC)
	fmt.Printf("Fells like: %.1f°C\n", forecast.Current.FellsLike)
	fmt.Printf("Precipitation: %.1fmm\n", forecast.Current.PrecipMM)
	fmt.Printf("Condition: %s\n", forecast.Current.Condition.Text)
	fmt.Printf("Humidity: %d%%\n", forecast.Current.Humidity)

	day := forecast.Forecast.Forecastday[0]
	for _, hour := range day.Hour {
		t := hour.Time
		formattedTime, err := time.Parse("2006-01-02 15:04", t)
		if err != nil {
			fmt.Println("Error parsing time:", err)
			return
		}
		// TODO it works but it should be improved
		if time.Now().Hour() <= formattedTime.Hour() {
			fmt.Printf("%s => %.1f°C, chance of rain: %.1f%%, will it rain: %d, %s\n", getTime(t), hour.TempC, hour.ChanceOfRain, hour.WillItRain, strings.Trim(hour.Condition.Text, " "))
		}

	}
}
