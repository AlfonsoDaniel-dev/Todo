package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"todoApp-backend/src/internal/domain"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db: db}
}

func (UR *userRepository) Save(user *domain.User) error {
	tx, err := UR.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(saveUserOnDB)
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

func (UR *userRepository) GetUserTask(id uuid.UUID) ([]domain.Task, error) {
	if id == uuid.Nil {
		return []domain.Task{}, errors.New(fmt.Sprintf("userRepository.GetUserTask: uuid cannot be nil"))
	}

	tx, err := UR.db.Begin()
	if err != nil {
		return []domain.Task{}, err
	}

	stmt, err := tx.Prepare(userGetItsTasks)
	if err != nil {
		return []domain.Task{}, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return []domain.Task{}, err
	}

	var tasks []domain.Task

	for rows.Next() {
		var task domain.Task

		err := rows.Scan(&task.Id, &task.OwnerId, &task.Title, &task.Body, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return []domain.Task{}, err
		}

		tasks = append(tasks, task)

	}

	if err := rows.Err(); err != nil {
		return []domain.Task{}, err
	}

	defer rows.Close()

	return tasks, tx.Commit()
}

func (UR *userRepository) GetUserData(email string) (domain.User, error) {
	if email == "" {
		return domain.User{}, errors.New("email is empty")
	}

	tx, err := UR.db.Begin()
	if err != nil {
		return domain.User{}, err
	}

	stmt, err := tx.Prepare(getUserOnDB)
	if err != nil {
		return domain.User{}, err
	}

	defer stmt.Close()

	var user domain.User

	err = stmt.QueryRow(email).Scan(&user)
	if err != nil {
		return domain.User{}, err
	}

	return user, tx.Commit()
}

func (UR *userRepository) GetIdByEmail(email string) (uuid.UUID, error) {
	if email == "" {
		return uuid.UUID{}, errors.New("email is empty")
	}

	tx, err := UR.db.Begin()
	if err != nil {
		return uuid.UUID{}, err
	}

	stmt, err := tx.Prepare(getUserIdByEmailOnDB)
	if err != nil {
		return uuid.UUID{}, err
	}

	defer stmt.Close()

	var id uuid.UUID

	err = stmt.QueryRow(email).Scan(&id)
	if err != nil {
		return uuid.UUID{}, err
	}

	return id, tx.Commit()
}

func (UR *userRepository) GetEmailById(id uuid.UUID) (string, error) {
	if id == uuid.Nil {
		return "", errors.New("id is empty")
	}

	tx, err := UR.db.Begin()
	if err != nil {
		return "", err
	}

	stmt, err := tx.Prepare(getUserEmailByIdOnDB)
	if err != nil {
		return "", err
	}

	defer stmt.Close()

	var email string
	err = stmt.QueryRow(id).Scan(&email)
	if err != nil {
		return "", err
	}

	return email, tx.Commit()
}
