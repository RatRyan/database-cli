package commands

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/RatRyan/dbapp/internal/employee"
	"github.com/RatRyan/dbapp/internal/paths"
	"github.com/RatRyan/dbapp/serialize"
	"google.golang.org/protobuf/proto"
)

func deserialize(id int) {
	data, err := os.ReadFile(paths.LongSer + strconv.Itoa(id) + ".protobuf")
	if err != nil {
		log.Fatal("File doesn't exist")
	}

	protoEmployee := serialize.Employee{}
	err = proto.Unmarshal(data, &protoEmployee)
	if err != nil {
		log.Fatal("wait protobuf error??? google wtf!")
	}

	employee := employee.Employee{
		Id:        int(protoEmployee.GetId()),
		FirstName: protoEmployee.GetFirstName(),
		LastName:  protoEmployee.GetLastName(),
		HireDate:  int(protoEmployee.GetHireDate()),
	}

	fmt.Println(employee)
}
