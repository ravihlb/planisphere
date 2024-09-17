package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(week)
}

var week = &cobra.Command{
	Use:   "week",
	Short: "Prints out the weekly plan",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Printing week plan...")
	},
}
