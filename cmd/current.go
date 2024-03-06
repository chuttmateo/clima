/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

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
	fmt.Printf("%s, %s, %s\n", forecast.Location.Name, forecast.Location.Region, forecast.Location.Country)

	fmt.Printf("Temperature: %.1f°C\n", forecast.Current.TempC)
	fmt.Printf("Precipitation: %.1fmm\n", forecast.Current.PrecipMM)
	fmt.Printf("Condition: %s\n", forecast.Current.Condition.Text)
	fmt.Printf("Fells like: %.1f°C\n", forecast.Current.FellsLike)
	fmt.Printf("Humidity: %d%%\n", forecast.Current.Humidity)
}
