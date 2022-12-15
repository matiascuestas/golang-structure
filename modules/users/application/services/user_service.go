package services

import (
	"github.com/example/modak/modules/users/application/services/interfaces"
	user_dtos "github.com/example/modak/modules/users/models/dtos"
	user_repo_interfaces "github.com/example/modak/modules/users/models/repositories/interfaces"
)

type userService struct {
	userRepository  user_repo_interfaces.IUserRepository
	rolesRepository user_repo_interfaces.IRolesRepository
}

func (s userService) GetUserById(id int) user_dtos.UserDTO {
	return s.userRepository.GetUserById(id)

}

func (s userService) SaveNewUser(user user_dtos.UserDTO) {
	s.rolesRepository.VerifyPermissions()
	s.userRepository.SaveNewUser(user)
}

func NewUserService(userRepository user_repo_interfaces.IUserRepository, rolesRepository user_repo_interfaces.IRolesRepository) interfaces.IUserService {
	return userService{
		userRepository:  userRepository,
		rolesRepository: rolesRepository,
	}
}
