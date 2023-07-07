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

var path = "C:/Users/rratajczak/Document/Neumont/2022-2023/Q4/DBT230/DatabaseProject/people"

func AddEmployee(id int, firstName string, lastName string, hireDate int) {
	content := []byte(strconv.Itoa(id) + ", " + firstName + ", " + lastName + strconv.Itoa(hireDate))
	err := os.WriteFile(path+strconv.Itoa(id)+".txt", content, 0777)
	if err != nil {
		panic(err)
	}
	fmt.Println("Employee successfully added!")
}

func DeleteEmployee(id int) {

}

func UpdateEmployee(id int, firstName string, lastName string, hireDate int) {

}
