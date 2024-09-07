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
	Short: "The current weather conditions",
	Long:  `This command will return the current weather conditions.`,
	Run: func(cmd *cobra.Command, args []string) {

		forecast, err := api.GetForecast(location, lang)
		if err != nil {
			fmt.Println("Error getting current forecast:", err)
			return
		}
		if emojis {
			printCurrentWithEmojis(forecast)
		} else {
			printCurrent(forecast)
		}
	},
}

func init() {
	rootCmd.AddCommand(currentCmd)
}

func printCurrent(forecast api.Forecast) {
	fmt.Printf(" %s - %s, %s, %s\n", forecast.Current.Condition.Text, forecast.Location.Name, forecast.Location.Region, forecast.Location.Country)
	fmt.Printf(" %.1f°C Fells like: %.1f°C\n", forecast.Current.TempC, forecast.Current.FellsLike)
	fmt.Printf(" Humidity: %d%%\n", forecast.Current.Humidity)

	//add a space
	fmt.Println()

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
			fmt.Printf(" %s => %.1f°C, chance of rain: %.1f%%, %s\n",
				getTime(t), hour.TempC, hour.ChanceOfRain, strings.Trim(hour.Condition.Text, " "))
		}

	}
}
func printCurrentWithEmojis(forecast api.Forecast) {

	fmt.Printf(" %s - %s, %s, %s\n", forecast.Current.Condition.Text, forecast.Location.Name, forecast.Location.Region, forecast.Location.Country)
	fmt.Printf(" %.1f°C Fells like: %.1f°C\n", forecast.Current.TempC, forecast.Current.FellsLike)
	fmt.Printf(" Humidity: %d%%\n", forecast.Current.Humidity)

	//add a space
	fmt.Println()

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
			fmt.Printf(" %s  %s => %.1f°C, chance of rain: %.1f%%, %s\n",
				weatherEmoji(hour.WillItRain), getTime(t), hour.TempC, hour.ChanceOfRain, strings.Trim(hour.Condition.Text, " "))
		}

	}
}

// chance of rain is an int8. 1 for true and 0 for false. returns a cloud if it is 1 otherwise returns a sun
func weatherEmoji(chance int8) string {
	if chance == 1 {
		return "\U000026C8\U0000FE0F"
	} else {
		return "\U00002600\U0000FE0F"
	}
}
