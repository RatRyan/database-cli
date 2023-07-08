package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/RatRyan/datacli/internal/paths"
	"github.com/RatRyan/datacli/serialize"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
)

var serializeAllCmd = &cobra.Command{
	Use:   "serializeAll",
	Short: "Serializes all of the files in a directory",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now()
		dir, err := os.ReadDir(paths.Long)
		if err != nil {
			log.Fatal("Directory doesn't exist")
		}

		for _, file := range dir {
			rawData, _ := os.ReadFile(paths.Long + file.Name())
			data := strings.Split(string(rawData), ", ")
			id, _ := strconv.Atoi(data[0])
			hireDate, _ := strconv.Atoi(data[3])
			employee := &serialize.Employee{
				Id:        int64(id),
				FirstName: data[1],
				LastName:  data[2],
				HireDate:  int64(hireDate),
			}
			message, err := proto.Marshal(employee)
			if err != nil {
				log.Fatal("wait protobuf error??? google wtf!")
			}
			os.WriteFile(paths.LongSer+data[0]+".ser", message, 0777)
		}
		fmt.Println("Finished Serialization! Time Elapsed:", time.Since(start))
	},
}

func init() {
	rootCmd.AddCommand(serializeAllCmd)
}
