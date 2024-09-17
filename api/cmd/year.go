package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(year)
}

var year = &cobra.Command{
	Use:   "year",
	Short: "Prints out the yearly plan",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Printing year plan...")
	},
}
