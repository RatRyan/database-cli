package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"datacli/internal/employee"
)

func printPeopleDetails(path string) {
	dir, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, file := range dir {
		content, _ := os.ReadFile(path + "/" + file.Name())
		fmt.Print(string(content))
	}
	fmt.Println()
}

func printEmployees(path string) {
	dir, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, file := range dir {
		rawData, _ := os.ReadFile(path + "/" + file.Name())
		trimData := strings.Trim(string(rawData), "\r\n")
		data := strings.Split(string(trimData), ", ")
		employee := new(internal.Employee)
		employee.Id, _ = strconv.Atoi(data[0])
		employee.FirstName = data[1]
		employee.LastName = data[2]
		employee.HireYear, _ = strconv.Atoi(data[3])
		fmt.Println(employee)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter path to directory: ")
	path, _ := reader.ReadString('\n')

	path = strings.TrimSpace(path)
	path = filepath.ToSlash(path)
	fmt.Println()

	printPeopleDetails(path)
	printEmployees(path)

	fmt.Print("Press enter to exit...")
	reader.ReadBytes('\n')
}
