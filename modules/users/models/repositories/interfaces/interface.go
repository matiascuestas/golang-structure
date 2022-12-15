package interfaces

import user_dtos "github.com/example/modak/modules/users/models/dtos"

type IUserRepository interface {
	GetUserById(int) user_dtos.UserDTO
	SaveNewUser(user_dtos.UserDTO)
}

type IPermissionRepository interface {
	Example()
}

type IRolesRepository interface {
	VerifyPermissions()
}
