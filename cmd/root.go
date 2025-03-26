package cmd

import "github.com/spf13/cobra"


func init() {
	rootCmd.AddCommand(versionCmd)
}


var rootCmd = &cobra.Command{
	Use: "expense-tracker ",
	Short: "Expense Tracker is a CLI application to track your expenses",
	Long: "Expense Tracker is a CLI application to track your expenses",
	Run: func(cmd *cobra.Command, args []string) {

	},
}