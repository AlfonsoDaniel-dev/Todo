package data

import (
	"database/sql"
	"github.com/google/uuid"
	"todoApp-backend/src/Core/domain"
)

type taskRepository struct {
	db *sql.DB
}

func (t *taskRepository) Save(task *domain.Task) error {
	if task == nil || task.Title == "" || task.Body == "" || task.OwnerId == uuid.Nil {
		return domain.ErrNotEnoughOrValidArguments
	}

	tx, err := t.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(saveTaskOnDB)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(task.Id, task.OwnerId, task.Title, task.Body, task.CreatedAt, task.UpdatedAt.Unix(), IntToNull(task.UpdatedAt.Unix()))
	if err != nil {
		return err
	}

	if err := stmt.Close(); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func NewTaskRepository(db *sql.DB) *taskRepository {
	return &taskRepository{
		db: db,
	}
}
