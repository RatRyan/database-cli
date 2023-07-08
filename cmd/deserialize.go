package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/RatRyan/datacli/internal/employee"
	"github.com/RatRyan/datacli/internal/paths"
	"github.com/RatRyan/datacli/serialize"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
)

var deserializeCmd = &cobra.Command{
	Use:   "deserialize",
	Short: "deserializes a serialized file",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		data, err := os.ReadFile(paths.LongSer + args[0] + ".protobuf")
		if err != nil {
			log.Fatal("File doesn't exist")
		}

		protoEmployee := serialize.Employee{}
		err = proto.Unmarshal(data, &protoEmployee)
		if err != nil {
			log.Fatal("wait protobuf error??? google wtf!")
		}

		employee := employee.Employee{
			Id:        int(protoEmployee.GetId()),
			FirstName: protoEmployee.GetFirstName(),
			LastName:  protoEmployee.GetLastName(),
			HireDate:  int(protoEmployee.GetHireDate()),
		}

		fmt.Println(employee)
	},
}

func init() {
	rootCmd.AddCommand(deserializeCmd)
}
