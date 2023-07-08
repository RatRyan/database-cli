package employee

import (
	"fmt"
	"os"
	"strconv"

	"github.com/RatRyan/dbapp/internal/paths"
)

func update(id int, firstName string, lastName string, hireDate int) {
	_, err := os.ReadFile(paths.Long + strconv.Itoa(id) + ".txt")
	if err != nil {
		panic(err)
	}
	WriteEmployee(id, firstName, lastName, hireDate)
	fmt.Println("Employee " + strconv.Itoa(id) + " Successfully updated")
}
