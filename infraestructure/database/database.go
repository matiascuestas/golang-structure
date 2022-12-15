package database

type IAuditable interface {
	GetUUID() string
}

type DataAccess interface {
	ReadAll() []IAuditable
	ReadById(int) IAuditable
	Save(IAuditable) bool
}
