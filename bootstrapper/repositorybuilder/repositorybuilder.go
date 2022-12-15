package repositorybuilder

import (
	"github.com/example/modak/bootstrapper/infrabuilder"
	"github.com/example/modak/modules/users/models/repositories"
	"github.com/example/modak/modules/users/models/repositories/interfaces"
)

var cacheUserRepository interfaces.IUserRepository
var cacheRolesRepository interfaces.IRolesRepository

func GetUserRepository() interfaces.IUserRepository {
	if cacheUserRepository == nil {
		cacheUserRepository = repositories.NewUserRepository(infrabuilder.GetDynamoDB())
	}
	return cacheUserRepository
}

func GetRolesRepository() interfaces.IRolesRepository {
	if cacheRolesRepository == nil {
		cacheUserRepository = repositories.NewUserRepository(infrabuilder.GetDynamoDB())
	}
	return cacheRolesRepository
}
