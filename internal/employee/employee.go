package employee

import (
	"fmt"
	"os"
	"strconv"
)

type Employee struct {
	Id        int
	FirstName string
	LastName  string
	HireYear  int
}

func (e Employee) String() string {
	return "ID:        " + strconv.Itoa(e.Id) + "\nfirstName: " + e.FirstName +
		"\nlastName:  " + e.LastName + "\nhireYear:  " + strconv.Itoa(e.HireYear) + "\n"
}

var path = "people/long/"

func writeEmployee(id int, firstName string, lastName string, hireDate int) {
	content := []byte(strconv.Itoa(id) + ", " + firstName + ", " + lastName + ", " + strconv.Itoa(hireDate))
	err := os.WriteFile(path + strconv.Itoa(id)+".txt", content, 0777)
	if err != nil {
		panic(err)
	}
}

func AddEmployee(id int, firstName string, lastName string, hireDate int) {
	writeEmployee(id, firstName, lastName, hireDate)
	fmt.Println("Employee successfully added")
}

func UpdateEmployee(id int, firstName string, lastName string, hireDate int) {
	_, err := os.ReadFile(path + strconv.Itoa(id) + ".txt")
	if err != nil {
		panic(err)
	}
	writeEmployee(id, firstName, lastName, hireDate)
	fmt.Println("Employee " + strconv.Itoa(id) + " Successfully updated")
}

func DeleteEmployee(id int) {
	err := os.Remove(path + strconv.Itoa(id) + ".txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Employee " + strconv.Itoa(id) + " successfully deleted")
}

