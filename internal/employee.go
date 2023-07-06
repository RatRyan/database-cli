package internal

import "strconv"

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