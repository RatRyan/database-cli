package commands

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/RatRyan/dbapp/internal/paths"
	"github.com/RatRyan/dbapp/serialize"
	"google.golang.org/protobuf/proto"
)

func serializeAll() {
	start := time.Now()
	dir, err := os.ReadDir(paths.Long)
	if err != nil {
		log.Fatal("Directory doesn't exist")
	}

	for _, file := range dir {
		rawData, _ := os.ReadFile(paths.Long + file.Name())
		data := strings.Split(string(rawData), ", ")
		id, _ := strconv.Atoi(data[0])
		hireDate, _ := strconv.Atoi(data[3])
		employee := &serialize.Employee{
			Id:        int64(id),
			FirstName: data[1],
			LastName:  data[2],
			HireDate:  int64(hireDate),
		}
		message, err := proto.Marshal(employee)
		if err != nil {
			log.Fatal("wait protobuf error??? google wtf!")
		}
		os.WriteFile(paths.LongSer+data[0]+".ser", message, 0777)
	}
	fmt.Println("Finished Serialization! Time Elapsed:", time.Since(start))
}
