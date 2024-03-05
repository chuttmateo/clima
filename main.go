package main

import (
	"github.com/spf13/cobra"

	"clima/commands"
)

func main() {
	var rootCmd = &cobra.Command{Use: "clima"}
	rootCmd.AddCommand(commands.CmdPrint)
	rootCmd.AddCommand(commands.CmdWeather)
	rootCmd.Execute()
}
