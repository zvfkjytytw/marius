package mariuspg

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

const (
	dataParts = 6
)

func (pg *PGAgent) GetMulus(fileID int32) ([dataParts]string, error) {
	sql, args, err := sq.Select(foldersTable).Columns("part1", "part2", "part3", "part4", "part5", "part6").
		Where(sq.Eq{"id": fileID}).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return [dataParts]string{}, fmt.Errorf("fail to generate select parts query for file %d: %w", fileID, err)
	}

	row := pg.db.QueryRow(sql, args...)
	var p1, p2, p3, p4, p5, p6 string
	err = row.Scan(&p1, &p2, &p3, &p4, &p5, &p6)
	if err != nil {
		return [dataParts]string{}, fmt.Errorf("file %d not found", fileID)
	}

	return [dataParts]string{p1, p2, p3, p4, p5, p6}, nil
}

func (pg *PGAgent) SetMulus(fileID int32, mulus [dataParts]string) error {
	if !pg.checkFile(fileID) {
		return fmt.Errorf("file %d not found", fileID)
	}

	sql, args, err := sq.Update(foldersTable).
		Set("part1", mulus[0]).Set("part2", mulus[1]).Set("part3", mulus[2]).
		Set("part4", mulus[3]).Set("part5", mulus[4]).Set("part6", mulus[5]).
		Where(sq.Eq{"id": fileID}).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return fmt.Errorf("fail to generate set mulus query for the file %d: %w", fileID, err)
	}

	tx, err := pg.db.Begin()
	if err != nil {
		return fmt.Errorf("fail start transaction: %w", err)
	}
	defer tx.Rollback()

	result, err := tx.Exec(sql, args...)
	if err != nil {
		return fmt.Errorf("fail to execute set mulus query for the file %d: %w", fileID, err)
	}

	if n, _ := result.RowsAffected(); n > 1 {
		return fmt.Errorf("affected %d rows instead 1", n)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("fail commit set mulus query: %w", err)
	}

	return nil
}

func (pg *PGAgent) FlushMulus(fileID int32) error {
	if !pg.checkFile(fileID) {
		return fmt.Errorf("file %d not found", fileID)
	}

	var empty string
	sql, args, err := sq.Update(foldersTable).
		Set("part1", empty).Set("part2", empty).Set("part3", empty).
		Set("part4", empty).Set("part5", empty).Set("part6", empty).
		Where(sq.Eq{"id": fileID}).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return fmt.Errorf("fail to generate flush mulus query for the file %d: %w", fileID, err)
	}

	tx, err := pg.db.Begin()
	if err != nil {
		return fmt.Errorf("fail start transaction: %w", err)
	}
	defer tx.Rollback()

	result, err := tx.Exec(sql, args...)
	if err != nil {
		return fmt.Errorf("fail to execute flush mulus query for the file %d: %w", fileID, err)
	}

	if n, _ := result.RowsAffected(); n > 1 {
		return fmt.Errorf("affected %d rows instead 1", n)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("fail commit fush mulus query: %w", err)
	}

	return nil
}

func (pg *PGAgent) checkFile(fileID int32) bool {
	sql, args, err := sq.Select(foldersTable).Columns("id").
		Where(sq.Eq{"id": fileID}).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return false
	}

	row := pg.db.QueryRow(sql, args...)
	var id int32
	if err := row.Scan(&id); err != nil {
		return false
	}

	return true
}
