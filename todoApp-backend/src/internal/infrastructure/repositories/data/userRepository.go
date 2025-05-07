package data

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"todoApp-backend/src/internal/domain"
	"todoApp-backend/src/internal/infrastructure/repositories"
)

type userRepository struct {
	db *sql.DB
}

func (UR *userRepository) GetUserData(id uuid.UUID) (*domain.User, error) {
	//TODO implement me

	if id == uuid.Nil {
		return nil, domain.ErrIdIsNotValid
	}

	stmt, err := UR.db.Prepare(repositories.getUserOnDB)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	user := &domain.User{}
	row := stmt.QueryRow(id)

	err = row.Scan(user.Id, user.Username, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, domain.ErrNotFound
	} else if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	return user, nil
}

func (UR *userRepository) GetIdByEmail(email string) (uuid.UUID, error) {
	if email == "" {
		return uuid.Nil, domain.ErrIdIsNotValid
	}

	stmt, err := UR.db.Prepare(repositories.getUserIdByEmailOnDB)
	if err != nil {
		return uuid.Nil, err
	}

	row := stmt.QueryRow(email)

	var id uuid.UUID
	err = row.Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		return uuid.Nil, domain.ErrNotFound
	} else if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return uuid.Nil, err
	}

	return id, nil
}

func (UR *userRepository) GetEmailById(id uuid.UUID) (string, error) {
	if id == uuid.Nil {
		return "", domain.ErrIdIsNotValid
	}

	stmt, err := UR.db.Prepare(repositories.getUserEmailByIdOnDB)
	if err != nil {
		return "", err
	}

	row := stmt.QueryRow(id)
	var email string

	err = row.Scan(&email)
	if errors.Is(err, sql.ErrNoRows) {
		return "", domain.ErrNotFound
	} else if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return "", err
	}

	return email, nil
}

func (UR *userRepository) GetUserPassword(id uuid.UUID) (string, error) {
	if id == uuid.Nil {
		return "", domain.ErrIdIsNotValid
	}
	stmt, err := UR.db.Prepare(repositories.getUserPassword)
	if err != nil {
		return "", err
	}

	row := stmt.QueryRow(id)
	var password string
	err = row.Scan(&password)
	if errors.Is(err, sql.ErrNoRows) {
		return "", domain.ErrNotFound
	} else if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return "", err
	}

	return password, nil
}

func (UR *userRepository) UpdateName(NewName string, id uuid.UUID) error {

	if NewName == "" || id == uuid.Nil {
		return domain.ErrNotEnoughOrValidArguments
	}

	tx, err := UR.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(repositories.userUpdateName)
	if err != nil {
		tx.Rollback()
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(NewName, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (UR *userRepository) UpdateEmail(NewEmail string, id uuid.UUID) error {
	if NewEmail == "" || id == uuid.Nil {
		return domain.ErrNotEnoughOrValidArguments
	}

	tx, err := UR.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(repositories.userUpdateEmail)
	if err != nil {
		tx.Rollback()
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(NewEmail, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (UR *userRepository) UpdatePassword(Password string, id uuid.UUID) error {

	if Password == "" || id == uuid.Nil {
		return domain.ErrNotEnoughOrValidArguments
	}

	tx, err := UR.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(repositories.userUpdatePassword)
	if err != nil {
		tx.Rollback()
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(Password, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (UR *userRepository) DeleteUser(id uuid.UUID) error {
	if id == uuid.Nil {
		return domain.ErrIdIsNotValid
	}

	tx, err := UR.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(repositories.userDeleteInfo)
	if err != nil {
		tx.Rollback()
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (UR *userRepository) CheckUserExists(email string) (bool, error) {
	if email == "" {
		return false, domain.ErrNotEnoughOrValidArguments
	}

	stmt, err := UR.db.Prepare(CheckUserExists)
	if err != nil {
		return false, err
	}

	row := stmt.QueryRow(email)
	var exists bool
	err = row.Scan(&exists)

	if !exists {
		return false, domain.ErrNotFound
	}

	return exists, nil
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db: db}
}

func (UR *userRepository) Save(user *domain.User) error {
	tx, err := UR.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(repositories.saveUserOnDB)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	return tx.Commit()
}
