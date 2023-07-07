package cmd

import (
	"github.com/spf13/cobra"
)

var employeeCmd = &cobra.Command{
	Use:   "employee",
	Short: "provides tools for managing employees",
	Long:  ``,
}

func init() {
	rootCmd.AddCommand(employeeCmd)
}
