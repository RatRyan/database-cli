package employee

import (
	"os"
	"strconv"
)

var Path = "C:/Users/rratajczak/Documents/Neumont/2022-2023/Q4/DBT230/datacli/people/long/"

type Employee struct {
	Id        int
	FirstName string
	LastName  string
	HireDate  int
}

func (e Employee) String() string {
	return "ID:        " + strconv.Itoa(e.Id) + "\nfirstName: " + e.FirstName +
		"\nlastName:  " + e.LastName + "\nhireDate:  " + strconv.Itoa(e.HireDate) + "\n"
}

func WriteEmployee(id int, firstName string, lastName string, hireDate int) {
	content := []byte(strconv.Itoa(id) + ", " + firstName + ", " + lastName + ", " + strconv.Itoa(hireDate))
	err := os.WriteFile(Path + strconv.Itoa(id)+".txt", content, 0777)
	if err != nil {
		panic(err)
	}
}