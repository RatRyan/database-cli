package main

import (
	"context"
	"time"

	"github.com/RatRyan/dbapp/internal/record/restaurant"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientOptions = options.Client().ApplyURI("mongodb://localhost:2717")
var client, err = mongo.Connect(context.Background(), clientOptions)

var database = client.Database("restaurantdb")
var collection = database.Collection("restaurants")

func CreateRecord(record interface{}, collectionName string) {
	collection = database.Collection(collectionName)
	collection.InsertOne(context.TODO(), record)
}

func readRecord() {

}

func updateRecord() {

}

func deleteRecord() {

}

func main() {
	defer client.Disconnect(context.Background())

	record := restaurant.Restaurant{
		Address: restaurant.AddressInfo{
			Building: "test building",
			Coord:    []float64{0.12, 3.45},
			Street:   "test street",
			Zipcode:  "test zipcode",
		},
		Borough: "test borough",
		Cuisine: "test cuisine",
		Grades: []restaurant.GradeInfo{{
			Date:  time.Now(),
			Grade: "A",
			Score: 5,
		}},
		Name:         "test name",
		RestaurantId: "test restaurant_id",
	}

	CreateRecord(record, "restaurants")
}
