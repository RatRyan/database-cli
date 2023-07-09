package console

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/RatRyan/dbapp/internal/employee"
	"github.com/RatRyan/dbapp/internal/serialize"
	"github.com/urfave/cli/v2"
	"google.golang.org/protobuf/proto"
)

var Running = true
var dirPath = "C:/Users/rratajczak/Documents/people/long"

var commands = []*cli.Command{
	{
		Name:        "path",
		Description: "Prints out the current path being used",
		Action: func(cCtx *cli.Context) error {
			fmt.Println("current path: " + dirPath)
			return nil
		},
		Subcommands: []*cli.Command{
			{
				Name:        "set",
				Description: "Sets the path to the people folder",
				Action: func(cCtx *cli.Context) error {
					dirPath = filepath.ToSlash(cCtx.Args().Get(0))
					fmt.Println("Path set to: '" + dirPath + "'")
					return nil
				},
			},
		},
	},
	{
		Name: "employee",
		Subcommands: []*cli.Command{
			{
				Name: "add",
				Action: func(cCtx *cli.Context) error {
					args := cCtx.Args().Slice()
					id, _ := strconv.Atoi(args[0])
					hireDate, _ := strconv.Atoi(args[3])
					employee.WriteEmployee(id, args[1], args[2], hireDate, dirPath+"/")
					fmt.Println("Employee successfully added")
					return nil
				},
			},
			{
				Name: "update",
				Action: func(cCtx *cli.Context) error {
					args := cCtx.Args().Slice()
					id, _ := strconv.Atoi(args[0])
					hireDate, _ := strconv.Atoi(args[3])
					_, err := os.ReadFile(dirPath + "/" + strconv.Itoa(id) + ".txt")
					if err != nil {
						log.Fatal("File doesn't exist")
					}
					employee.WriteEmployee(id, args[1], args[2], hireDate, dirPath+"/")
					fmt.Println("Employee " + strconv.Itoa(id) + " Successfully updated")
					return nil
				},
			},
			{
				Name: "delete",
				Action: func(cCtx *cli.Context) error {
					err := os.Remove(dirPath + "/" + cCtx.Args().First() + ".txt")
					if err != nil {
						panic(err)
					}
					fmt.Println("Employee " + cCtx.Args().First() + " successfully deleted")
					return nil
				},
			},
		},
	},
	{
		Name:        "serialize",
		Description: "serializes all of the files in a directory",
		Action: func(cCtx *cli.Context) error {

			dir, err := os.ReadDir(dirPath)
			if err != nil {
				log.Fatal("Directory doesn't exist")
			}

			var wg sync.WaitGroup
			start := time.Now()
			for _, file := range dir {
				wg.Add(1)

				go func(entry os.DirEntry) {
					rawData, _ := os.ReadFile(filepath.Join(dirPath, entry.Name()))
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
						log.Fatal("protobuf error")
					}
					os.WriteFile(filepath.Join(dirPath+" serialized", data[0]+".ser"), message, 0777)
					wg.Done()
				}(file)
			}

			fmt.Println("Finished Serialization! Time Elapsed:", time.Since(start))
			return nil
		},
	},
	{
		Name:        "deserialize",
		Description: "deserialize a file given the id",
		Action: func(cCtx *cli.Context) error {
			data, err := os.ReadFile(dirPath + " serialized/" + cCtx.Args().Get(0) + ".ser")
			if err != nil {
				log.Fatal("File doesn't exist")
			}

			protoEmployee := serialize.Employee{}
			err = proto.Unmarshal(data, &protoEmployee)
			if err != nil {
				log.Fatal("protobuf error")
			}

			employee := employee.Employee{
				Id:        int(protoEmployee.GetId()),
				FirstName: protoEmployee.GetFirstName(),
				LastName:  protoEmployee.GetLastName(),
				HireDate:  int(protoEmployee.GetHireDate()),
			}

			fmt.Println(employee)
			return nil
		},
	},
	{
		Name:        "quit",
		Description: "quits out of the application",
		Action: func(cCtx *cli.Context) error {
			Running = false
			return nil
		},
	},
}

var App = &cli.App{
	Name:        "Database tool",
	Description: "Console Application for reading/writing/serializing a .txt based database",
	Commands:    commands,
}
