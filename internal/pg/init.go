package mariuspg

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed init.sql
var initDB string

func (pg *PGAgent) initDB() error {
	tx, err := pg.db.Begin()
	if err != nil {
		return fmt.Errorf("failed create init DB transaction: %v", err)
	}
	defer tx.Rollback()

	initQuerysFS := strings.Split(initDB, "/n")

	for _, query := range initQuerysFS {
		if _, err = tx.Exec(query); err != nil {
			return fmt.Errorf("failed execute init DB query: %v", err)
		}
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed commit init DB querys: %v", err)
	}

	return nil
}
