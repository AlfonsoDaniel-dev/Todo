package user

import (
	"errors"
	"github.com/google/uuid"
	"strings"
	"todoApp-backend/src/internal/domain"
	"todoApp-backend/src/internal/infrastructure/DTO"
)

func (S *UserServices) CreateUser(UserDTO *DTO.UserDTO) error {
	if UserDTO.Username == "" || UserDTO.Email == "" || UserDTO.Password == "" {
		return errors.New("username or email or password is empty")
	}

	UserDTO.Password = strings.TrimSpace(UserDTO.Password)
	UserDTO.Email = strings.TrimSpace(UserDTO.Email)

	User, err := domain.NewUser(UserDTO.Username, UserDTO.Email, UserDTO.Password)
	if err != nil {
		return err
	}

	userExists, err := S.Repository.CheckUserExists(User.Id)
	if userExists {
		return errors.New("user already exists")
	}

	err = S.Repository.Save(&User)
	if err != nil {
		return err
	}

	return nil
}

func (S *UserServices) OAuthLogin(userName, email string) error {
	if userName == "" || email == "" {
		return errors.New("username or email are required")
	}

	id, err := S.Repository.GetIdByEmail(email)
	if id != uuid.Nil && err == nil {

		return domain.UserAlreadyExists

	} else if errors.Is(err, domain.ErrNotFound) {
		userPassword := domain.GeneratePassword()

		user, err := domain.NewUser(userName, email, userPassword)
		if err != nil {
			return err
		}

		err = S.Repository.Save(&user)
		if err != nil {
			return err
		}

		return domain.ErrNotFound
	}

	return err
}
