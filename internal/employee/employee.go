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
	HireDate  int
}

func (e Employee) String() string {
	return "ID:        " + strconv.Itoa(e.Id) + "\nfirstName: " + e.FirstName +
		"\nlastName:  " + e.LastName + "\nhireDate:  " + strconv.Itoa(e.HireDate) + "\n"
}

func WriteEmployee(id int, firstName string, lastName string, hireDate int, path string) {
	fmt.Println("'" + strconv.Itoa(hireDate) + "'")
	content := []byte(strconv.Itoa(id) + ", " + firstName + ", " + lastName + ", " + strconv.Itoa(hireDate))
	
	err := os.WriteFile(path+strconv.Itoa(id)+".txt", content, 0777)
	if err != nil {
		panic(err)
	}
}
