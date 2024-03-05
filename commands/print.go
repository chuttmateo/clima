package commands

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var CmdPrint = &cobra.Command{
	Use: "print",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(strings.Join(args, " "))
	},
}
