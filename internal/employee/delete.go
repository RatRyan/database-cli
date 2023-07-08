package employee

import (
	"fmt"
	"os"
	"strconv"

	"github.com/RatRyan/dbapp/internal/paths"
)

func Delete(id int) {
	err := os.Remove(paths.Long + strconv.Itoa(id) + ".txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("Employee " + strconv.Itoa(id) + " successfully deleted")
}
