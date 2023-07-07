package cmd

import (
	"fmt"
	"os"

	"github.com/RatRyan/datacli/internal/paths"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "deletes an existing employee",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := os.Remove(paths.Long + args[0] + ".txt")
		if err != nil {
			panic(err)
		}
		fmt.Println("Employee " + args[0] + " successfully deleted")
	},
}

func init() {
	employeeCmd.AddCommand(deleteCmd)
}
