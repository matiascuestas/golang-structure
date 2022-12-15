package servicebuilder

import (
	"github.com/example/modak/bootstrapper/repositorybuilder"
	"github.com/example/modak/infraestructure/database"
	"github.com/example/modak/infraestructure/database/dynamo"
	"github.com/example/modak/modules/users/application/services"
	userServiceInterface "github.com/example/modak/modules/users/application/services/interfaces"
	"github.com/example/modak/modules/users/models/dtos"
)

func init() {
	dynamo.ReadByIdMock = func(id int) database.IAuditable {
		user := dtos.UserDTO{
			UserName: "JuanPerez123",
		}
		return user
	}
}

var cacheUserService userServiceInterface.IUserService

func GetUserService() userServiceInterface.IUserService {
	if cacheUserService == nil {
		cacheUserService = services.NewUserService(repositorybuilder.GetUserRepository(), repositorybuilder.GetRolesRepository())
	}
	return cacheUserService
}
