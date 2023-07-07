package cmd

import (
	"fmt"
	"strconv"

	"github.com/RatRyan/datacli/internal/employee"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "creates a new employee",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := strconv.Atoi(args[0])
		firstName := args[1]
		lastName := args[2]
		hireDate, _ := strconv.Atoi(args[3])
		employee.WriteEmployee(id, firstName, lastName, hireDate)
		fmt.Println("Employee successfully added")
	},
}

func init() {
	employeeCmd.AddCommand(addCmd)
}
