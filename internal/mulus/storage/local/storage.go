package mulusstorage

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	Folder string `yaml:"folder"`
}

type Storage struct {
	folder string
}

func NewStorage(cfg Config) (*Storage, error) {
	absPath, err := filepath.Abs(cfg.Folder)
	if err != nil {
		return nil, err
	}

	if err := os.Mkdir(absPath, 0755); err != nil && !os.IsExist(err) {
		return nil, err
	}

	return &Storage{
		folder: absPath,
	}, nil
}

func (s *Storage) SaveData(name string, data []byte) error {
	filePath := filepath.Join(s.folder, name)
	if existFile(filePath) {
		return fmt.Errorf("file %s already exists", filePath)
	}

	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("failed write file %s: %w", filePath, err)
	}

	return nil
}

func (s *Storage) GetData(name string) ([]byte, error) {
	filePath := filepath.Join(s.folder, name)
	if !existFile(filePath) {
		return nil, fmt.Errorf("file %s not found", filePath)
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed read file %s: %w", filePath, err)
	}

	return data, nil
}

func (s *Storage) DeleteData(name string) error {
	filePath := filepath.Join(s.folder, name)
	if !existFile(filePath) {
		return fmt.Errorf("file %s not found", filePath)
	}

	if err := os.Remove(filePath); err != nil {
		return fmt.Errorf("failed remove file %s: %w", filePath, err)
	}

	return nil
}

func existFile(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}
