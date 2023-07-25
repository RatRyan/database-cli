package employee

type Employee struct {
	Id        int    `bson:"id"`
	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name"`
	HireDate  int    `bson:"hire_date"`
}
