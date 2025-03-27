package user

import (
	"errors"
	"strings"
	"todoApp-backend/src/internal/domain"
	"todoApp-backend/src/internal/infrastructure/controllers/DTO"
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

	userExists, err := S.Repository.CheckUserExists(UserDTO.Email)
	if userExists {
		return errors.New("user already exists")
	}

	err = S.Repository.Save(&User)
	if err != nil {
		return err
	}

	return nil
}

func (S *UserServices) UpdateUserName(form DTO.UpdateUserName, userEmail string) error {
	if form.NewUserName == "" || userEmail == "" {
		return errors.New("all fields are required")
	}

	userId, err := S.Repository.GetIdByEmail(userEmail)
	if err != nil {
		return err
	}

	err = S.Repository.UpdateName(form.NewUserName, userId)
	if err != nil {
		return err
	}

	return nil
}

func (S *UserServices) UpdateUserEmail(form DTO.UpdateUserEmail, userEmail string) error {
	if form.NewEmail == "" || userEmail == "" {
		return errors.New("all fields are required")
	}

	userId, err := S.Repository.GetIdByEmail(userEmail)
	if err != nil {
		return err
	}

	err = S.Repository.UpdateEmail(form.NewEmail, userId)
	if err != nil {
		return err
	}

	return nil
}

func (S *UserServices) UpdateUserPassword(form DTO.UpdateUserPassword, userEmail string) error {
	if form.NewPassword == "" || userEmail == "" {
		return errors.New("all fields are required")
	}

	userId, err := S.Repository.GetIdByEmail(userEmail)
	if err != nil {
		return err
	}

	oldPassword, err := S.Repository.GetUserPassword(userId)
	if err != nil {
		return err
	}

	ok := domain.ComparePassword(oldPassword, form.OldPassword)
	if !ok {
		return domain.ErrWrongPassword
	}

	err = S.Repository.UpdatePassword(form.NewPassword, userId)
	if err != nil {
		return err
	}

	return nil
}

func (S *UserServices) DeleteUser(deleteForm DTO.DeleteUser, email string) error {
	if deleteForm.Password == "" {
		return errors.New("old password is empty")
	}

	ok, err := S.Repository.CheckUserExists(email)
	if err != nil {
		return err
	} else if !ok {
		return domain.ErrNotFound
	}

	usrId, err := S.Repository.GetIdByEmail(email)
	if err != nil {
		return err
	}

	usrPassword, err := S.Repository.GetUserPassword(usrId)
	if err != nil {
		return err
	}
	ok = domain.ComparePassword(deleteForm.Password, usrPassword)
	if !ok {
		return domain.ErrWrongPassword
	}

	err = S.Repository.DeleteUser(usrId)
	if err != nil {
		return err
	}

	return nil
}
