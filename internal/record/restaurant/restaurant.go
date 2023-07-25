package restaurant

import "time"

type AddressInfo struct {
	Building string    `bson:"building"`
	Coord    []float64 `bson:"coord"`
	Street   string    `bson:"street"`
	Zipcode  string    `bson:"zipcode"`
}

type GradeInfo struct {
	Date  time.Time `bson:"date"`
	Grade string    `bson:"grade"`
	Score int32     `bson:"score"`
}

type Restaurant struct {
	Address      AddressInfo `bson:"address"`
	Borough      string      `bson:"borough"`
	Cuisine      string      `bson:"cuisine"`
	Grades       []GradeInfo `bson:"grades"`
	Name         string      `bson:"name"`
	RestaurantId string      `bson:"restaurant_id"`
}
