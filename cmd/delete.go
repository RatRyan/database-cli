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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
