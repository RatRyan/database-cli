package employee

import (
	"fmt"
)

func Add(id int, firstName string, lastName string, hireDate int) {
	WriteEmployee(id, firstName, lastName, hireDate)
	fmt.Println("Employee successfully added")
}
