package user

type User struct {
	DataInterface
}

func NewUser(dataInterface DataInterface) *User {
	return &User{
		dataInterface,
	}
}
