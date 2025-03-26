package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)


var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Print the version number of Expense Tracker",
	Long: "Print the version number of Expense Tracker",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("v0.1 -- HEAD")
	},
}