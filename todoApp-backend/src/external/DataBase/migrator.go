package DataBase

import (
	"database/sql"
	"log"
)

type Migrator struct {
	db *sql.DB
}

func (m *Migrator) Migrate() error {

	querys := []string{userTable, taskTable}

	tx, err := m.db.Begin()
	if err != nil {
		log.Fatalf("couldn't create transaction: %v", err)
	}

	for _, query := range querys {
		stmt, err := tx.Prepare(query)
		if err != nil {
			log.Fatalf("couldn't prepare statement: %v", err)
		}

		_, err = stmt.Exec()
		if err != nil {
			log.Fatalf("couldn't execute statement: %v", err)
		}
		if err := stmt.Close(); err != nil {
			log.Fatalf("couldn't close statement: %v", err)
		}
	}

	if err := tx.Commit(); err != nil {
		log.Fatalf("couldn't commit transaction: %v", err)
	}

	return nil
}

func NewMigrator(bd *sql.DB) *Migrator {
	return &Migrator{
		db: bd,
	}
}
