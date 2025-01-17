package gaiusstorage

import (
	"context"
	"fmt"
	"sync"

	"go.uber.org/zap"
)

type tessarius interface {
	GetMulus(int32) ([dataParts]string, error)
	SetMulus(int32, [dataParts]string) error
	FlushMulus(int32) error
}

type Storage struct {
	sync.RWMutex
	castra map[string]*legionarius
	cohors []*mus
	tesser tessarius
	logger *zap.Logger
}

func NewStorage(t tessarius, logger *zap.Logger) (*Storage, error) {
	return &Storage{
		castra: make(map[string]*legionarius),
		cohors: make([]*mus, 0, dataParts),
		tesser: t,
		logger: logger,
	}, nil
}

func (s *Storage) Kill(ctx context.Context) {
	for _, leg := range s.castra {
		leg.conn.Close()
	}
}

func (s *Storage) GetName() (name string, err error) {
	name = fmt.Sprintf("IS %s", "TEST")

	return
}
