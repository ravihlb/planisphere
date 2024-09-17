package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(day)
}

var day = &cobra.Command{
	Use:   "day",
	Short: "Prints out the daily plan",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Printing day plan...")
	},
}
