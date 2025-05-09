package DataBase

import (
	"database/sql"
	"log"
)

type Migrator struct {
	db *sql.DB
}

func (m *Migrator) Migrate() error {

	querys := []string{uuidExtension, userTable, taskTable}

	tx, err := m.db.Begin()
	if err != nil {
		log.Fatalf("couldn't create transaction: %v", err)
	}

	for i, query := range querys {
		stmt, err := tx.Prepare(query)
		if err != nil {
			if err := tx.Rollback(); err != nil {
				log.Fatalf("couldn't rollback transaction: %v", err)
			}
			log.Fatalf("couldn't prepare statement: %v failed on query n: %v", err, i)
		}

		_, err = stmt.Exec()
		if err != nil {
			if err := tx.Rollback(); err != nil {
				log.Fatalf("couldn't rollback transaction: %v", err)
			}
			log.Fatalf("couldn't execute statement: %v", err)
		}

		if err := stmt.Close(); err != nil {
			if err := tx.Rollback(); err != nil {
				log.Fatalf("couldn't rollback transaction: %v", err)
			}
			log.Fatalf("couldn't close statement: %v", err)
		}
	}

	if err := tx.Commit(); err != nil {
		if err := tx.Rollback(); err != nil {
			log.Fatalf("couldn't rollback transaction: %v", err)
		}
		log.Fatalf("couldn't commit transaction: %v", err)
	}

	return nil
}

func NewMigrator(bd *sql.DB) *Migrator {
	return &Migrator{
		db: bd,
	}
}
