package cmd

import (
	"github.com/spf13/cobra"
)

// employeeCmd represents the employee command
var employeeCmd = &cobra.Command{
	Use:   "employee",
	Short: "provides tools for managing employees",
	Long:  ``,
	// Run: func(cmd *cobra.Command, args []string) {
	// },
}

func init() {
	rootCmd.AddCommand(employeeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// employeeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// employeeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
