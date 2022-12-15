package repositories

import (
	"fmt"

	"github.com/example/modak/infraestructure/database"
	user_interfaces "github.com/example/modak/modules/users/models/repositories/interfaces"
)

type rolesRepository struct {
	db database.DataAccess
}

func (s rolesRepository) VerifyPermissions() {
	fmt.Println("call verifypermissions")
}

func NewRolesRepository(database database.DataAccess) user_interfaces.IRolesRepository {
	return &rolesRepository{db: database}
}
