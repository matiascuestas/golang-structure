package infrabuilder

import (
	"github.com/example/modak/infraestructure/database"
	"github.com/example/modak/infraestructure/database/dynamo"
)

var dynamoDB database.DataAccess

func GetDynamoDB() database.DataAccess {
	if dynamoDB == nil {
		dynamoDB = dynamo.NewDynamoDB("conn")
	}
	return dynamoDB
}
