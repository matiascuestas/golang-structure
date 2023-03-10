package dynamo

import (
	"fmt"

	"github.com/example/modak/infraestructure/database"
)

type dynamoDB struct {
	conn string
}

var ReadByIdMock func(id int) database.IAuditable

func (s dynamoDB) ReadAll() []database.IAuditable {
	fmt.Println("Readall")
	var collection []database.IAuditable
	return collection
}

func (s dynamoDB) ReadById(id int) database.IAuditable {
	fmt.Printf("id to get: %v \n", id)
	return ReadByIdMock(id)
}

func (s dynamoDB) Save(item database.IAuditable) bool {
	fmt.Printf("id to save: %v", item)
	return true
}

func NewDynamoDB(connectionString string) database.DataAccess {
	return dynamoDB{conn: connectionString}
}
