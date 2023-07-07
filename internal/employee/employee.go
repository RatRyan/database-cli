package employee

import (
	"strconv"
	"os"
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

func AddEmployee(id int, firstName string, lastName string, hireDate int) {
	os.WriteFile()
}

func DeleteEmployee(id int) {
	
}

func UpdateEmployee(id int, firstName string, lastName string, hireDate int) {

}