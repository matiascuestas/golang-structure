package repositorybuilder

import (
	"github.com/example/modak/infraestructure/database"
	"github.com/example/modak/infraestructure/database/dynamo"
	"github.com/example/modak/modules/users/models/repositories"
	"github.com/example/modak/modules/users/models/repositories/interfaces"
)

var dynamoDB database.DataAccess
var cacheUserRepository interfaces.IUserRepository
var cacheRolesRepository interfaces.IRolesRepository

func getDynamoDB() database.DataAccess {
	if dynamoDB == nil {
		dynamoDB = dynamo.NewDynamoDB("conn")
	}
	return dynamoDB
}

func GetUserRepository() interfaces.IUserRepository {
	if cacheUserRepository == nil {
		cacheUserRepository = repositories.NewUserRepository(getDynamoDB())
	}
	return cacheUserRepository
}

func GetRolesRepository() interfaces.IRolesRepository {
	if cacheRolesRepository == nil {
		cacheUserRepository = repositories.NewUserRepository(getDynamoDB())
	}
	return cacheRolesRepository
}
