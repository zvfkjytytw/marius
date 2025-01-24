package mariuspg

import (
	"errors"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
)

const (
	foldersTable       = "fs.folders"
	filesTable         = "fs.files"
	slash              = "/"
	rootFolderID int32 = 1
	rootUserID   int32 = 1
)

func (pg *PGAgent) FindFile(path string) (int32, error) {
	return pg.getItemID(path, filesTable)
}

func (pg *PGAgent) CreateFile(path string) (int32, error) {
	return pg.createFile(path, rootUserID)
}

func (pg *PGAgent) DeleteFile(path string) error {
	sql, args, err := sq.Delete(filesTable).Where(sq.Eq{"full_path": path}).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return fmt.Errorf("failed generate delete file %s query: %v", path, err)
	}

	_, err = pg.db.Exec(sql, args...)
	if err != nil {
		return fmt.Errorf("failed delete file %s from fs: %v", path, err)
	}

	return nil
}

func (pg *PGAgent) getItemID(path, itemTable string) (int32, error) {
	var itemID int32
	sql, args, err := sq.Select("id").From(itemTable).
		Where(sq.Eq{"full_path": path}).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return 0, fmt.Errorf("failed generate select ID query for folder %s: %v", path, err)
	}

	if err = pg.db.QueryRow(sql, args...).Scan(&itemID); err != nil {
		return 0, fmt.Errorf("item %s not found: %v", path, err)
	}

	return itemID, nil
}

func (pg *PGAgent) createFolder(path string, owner int32) (int32, error) {
	folderID, err := pg.getItemID(path, foldersTable)
	if err == nil {
		return folderID, nil
	}

	folder, parent := splitPath(path)
	parentID, err := pg.createFolder(parent, owner)
	if err != nil {
		return 0, fmt.Errorf("failed create folder %s in fs: %v", path, err)
	}

	sql, args, err := sq.Insert(foldersTable).Columns("name", "parent", "owner", "full_path").
		Values(folder, parentID, owner, path).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return 0, err
	}

	_, err = pg.db.Exec(sql, args...)
	if err != nil {
		return 0, err
	}

	return pg.getItemID(path, foldersTable)
}

func (pg *PGAgent) createFile(path string, owner int32) (int32, error) {
	if path == "" {
		return 0, errors.New("empty path")
	}

	fileName, folderPath := splitPath(path)
	parentID, err := pg.createFolder(folderPath, owner)
	if err != nil {
		return 0, fmt.Errorf("failed create parent folder for file %s: %v", path, err)
	}

	sql, args, err := sq.Insert(filesTable).Columns("name", "parent", "owner", "full_path").
		Values(fileName, parentID, owner, path).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return 0, fmt.Errorf("failed generate create file %s query: %v", path, err)
	}

	_, err = pg.db.Exec(sql, args...)
	if err != nil {
		return 0, fmt.Errorf("failed create file %s in fs: %v", path, err)
	}

	sql, args, err = sq.Select("id").From(filesTable).
		Where(sq.Eq{"name": fileName}).Where(sq.Eq{"parent": parentID}).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return 0, fmt.Errorf("failed generate get file %s ID query: %v", path, err)
	}

	var fileID int32
	if err = pg.db.QueryRow(sql, args...).Scan(&fileID); err != nil {
		return 0, fmt.Errorf("file %s not found: %v", fileName, err)
	}

	return fileID, nil
}

func splitPath(path string) (targetName, parentPath string) {
	pathSlice := strings.Split(path, slash)
	targetName = pathSlice[len(pathSlice)-1]
	parentPath = strings.Join(pathSlice[:len(pathSlice)-1], slash)

	return
}
