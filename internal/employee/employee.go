package employee

import (
	"os"
	"strconv"

	"github.com/RatRyan/dbapp/internal/paths"
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

func WriteEmployee(id int, firstName string, lastName string, hireDate int) {
	content := []byte(strconv.Itoa(id) + ", " + firstName + ", " + lastName + ", " + strconv.Itoa(hireDate))
	err := os.WriteFile(paths.Long + strconv.Itoa(id)+".txt", content, 0777)
	if err != nil {
		panic(err)
	}
}