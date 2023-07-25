package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/RatRyan/dbapp/internal/record/employee"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const PATH = "C:/Users/rratajczak/Documents/people/long/"

var clientOptions = options.Client().ApplyURI("mongodb://localhost:2717")
var client, err = mongo.Connect(context.Background(), clientOptions)
var database = client.Database("peopledb")
var collection = database.Collection("people")

func CreateRecord(record interface{}) {
	collection.InsertOne(context.TODO(), record)
}

func readRecord(id int) {
	var result employee.Employee
	filter := bson.D{{Key: "id", Value: id}}
	collection.FindOne(context.TODO(), filter).Decode(&result)
	fmt.Println(result)
}

func updateRecord(id int) {
	filter := bson.D{{Key: "id", Value: id}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "first_name", Value: "Chris Wilson"}}}}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
}

func deleteRecord(id int) {
	filter := bson.D{{Key: "id", Value: id}}
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
}

func initRecords() {
	dir, _ := os.ReadDir(PATH)

	var wg sync.WaitGroup
	for _, file := range dir {
		wg.Add(1)
		go func(entry os.DirEntry) {
			data, _ := os.ReadFile(PATH + entry.Name())
			employeeInfo := strings.Split(string(data), ", ")

			id, _ := strconv.Atoi(employeeInfo[0])
			hireDate, _ := strconv.Atoi(strings.TrimSuffix(employeeInfo[3], "\r\n"))

			record := employee.Employee{
				Id:        id,
				FirstName: employeeInfo[1],
				LastName:  employeeInfo[2],
				HireDate:  hireDate,
			}
			CreateRecord(record)
			wg.Done()
		}(file)
		wg.Wait()
	}
}

func main() {
	defer client.Disconnect(context.Background())

	deleteRecord(1)
}
