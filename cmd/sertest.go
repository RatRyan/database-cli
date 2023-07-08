/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/RatRyan/datacli/serialize"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
)

// sertestCmd represents the sertest command
var sertestCmd = &cobra.Command{
	Use:   "sertest",
	Short: "im gonna fucking cry",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		test := &serialize.Employee {
			Id:        0,
			FirstName: "Ryan",
			LastName:  "Ratajczak",
			HireDate:  2003,
		}

		fmt.Println("Marshal data:")
		data, err := proto.Marshal(test)
		if err != nil {
			panic(err)
		}
		println(data)

		fmt.Println("Unmarshal data:")
		employeeFromSer := serialize.Employee{}
		err = proto.Unmarshal(data, &employeeFromSer)
		if err != nil {
			panic(err)
		}
		fmt.Println(employeeFromSer.Id)
		fmt.Println(employeeFromSer.FirstName)
		fmt.Println(employeeFromSer.LastName)
		fmt.Println(employeeFromSer.HireDate)
	},
}

func init() {
	rootCmd.AddCommand(sertestCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sertestCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sertestCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
