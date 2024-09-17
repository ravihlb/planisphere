package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(month)
}

var month = &cobra.Command{
	Use:   "month",
	Short: "Prints out the monthly plan",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Printing month plan...")
	},
}
