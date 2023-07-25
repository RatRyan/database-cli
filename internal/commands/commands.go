package commands

// func deserialize(id string) employee.Employee {
// 	data, err := os.ReadFile(dirPath + " serialized/" + id + ".ser")
// 	if err != nil {
// 		log.Fatal("File doesn't exist: " + id + ".ser")
// 	}

// 	protoEmployee := serialize.Employee{}
// 	err = proto.Unmarshal(data, &protoEmployee)
// 	if err != nil {
// 		log.Fatal("protobuf error")
// 	}

// 	return employee.Employee{
// 		Id:        int(protoEmployee.GetId()),
// 		FirstName: protoEmployee.GetFirstName(),
// 		LastName:  protoEmployee.GetLastName(),
// 		HireDate:  int(protoEmployee.GetHireDate()),
// 	}
// }

// func doSerialize() {
// 	dir, err := os.ReadDir(dirPath)
// 	if err != nil {
// 		fmt.Println("Directory doesn't exist")
// 		return
// 	}

// 	var wg sync.WaitGroup
// 	start := time.Now()

// 	for _, file := range dir {
// 		wg.Add(1)
// 		go func(entry os.DirEntry) {
// 			rawData, _ := os.ReadFile(filepath.Join(dirPath, entry.Name()))
// 			data := strings.Split(strings.TrimSpace(string(rawData)), ", ")

// 			id, _ := strconv.Atoi(data[0])
// 			hireDate, _ := strconv.Atoi(data[3])

// 			employee := &serialize.Employee{
// 				Id:        int64(id),
// 				FirstName: data[1],
// 				LastName:  data[2],
// 				HireDate:  int64(hireDate),
// 			}

// 			message, err := proto.Marshal(employee)
// 			if err != nil {
// 				fmt.Println("Protobuf error")
// 			}

// 			err = os.WriteFile(filepath.Join(dirPath+" serialized", data[0]+".ser"), message, 0777)
// 			if err != nil {
// 				fmt.Println("File failed to write")
// 			}

// 			wg.Done()
// 		}(file)
// 	}
// 	fmt.Println("Finished Serialization! Time Elapsed:", time.Since(start))
// }
