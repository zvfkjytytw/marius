package gaiusstorage

import (
	"context"
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
	for addr, leg := range s.castra {
		if err := leg.conn.Close(); err != nil {
			s.logger.Sugar().Errorf("fail close connection to %s: %v", addr, err)
		}
	}
}
