package interfaces

import user_dtos "github.com/example/modak/modules/users/models/dtos"

type IUserService interface {
	GetUserById(int) user_dtos.UserDTO
	SaveNewUser(user_dtos.UserDTO)
}
