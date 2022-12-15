package dtos

type UserDTO struct {
	UUID       string
	Email      string
	UserName   string `json:"username"`
	Password   string
	CreateDate string
}

func (s UserDTO) GetUUID() string {
	return s.UUID
}
