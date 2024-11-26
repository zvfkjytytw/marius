package gaiusstorage

import (
	"fmt"
)

type Mus interface {
	DelData(string) error
	GetData(string) ([]byte, error)
	PutData(string, []byte) (string, error)
}

type GStorage struct {
	Mulus map[string]Mus
}

func (s *GStorage) GetName() (name string, err error) {
	name = fmt.Sprintf("IS %s", "TEST")

	return
}