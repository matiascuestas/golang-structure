package repositories

import (
	"github.com/example/modak/infraestructure/database"
	user_dtos "github.com/example/modak/modules/users/models/dtos"
	user_interfaces "github.com/example/modak/modules/users/models/repositories/interfaces"
)

type userRepository struct {
	db database.DataAccess
}

func (s userRepository) GetUserById(id int) user_dtos.UserDTO {
	auditable := s.db.ReadById(id)
	return auditable.(user_dtos.UserDTO)
}

func (s userRepository) SaveNewUser(user user_dtos.UserDTO) {
	s.db.Save(user)
}

func NewUserRepository(database database.DataAccess) user_interfaces.IUserRepository {
	return &userRepository{db: database}
}
