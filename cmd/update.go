package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/RatRyan/datacli/internal/employee"
	"github.com/RatRyan/datacli/internal/paths"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "updates an existing employee",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := strconv.Atoi(args[0])
		firstName := args[1]
		lastName := args[2]
		hireDate, _ := strconv.Atoi(args[3])
		_, err := os.ReadFile(paths.Long + strconv.Itoa(id) + ".txt")
		if err != nil {
			panic(err)
		}
		employee.WriteEmployee(id, firstName, lastName, hireDate)
		fmt.Println("Employee " + strconv.Itoa(id) + " Successfully updated")
	},
}

func init() {
	employeeCmd.AddCommand(updateCmd)
}
