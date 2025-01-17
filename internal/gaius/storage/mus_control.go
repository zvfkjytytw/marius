package gaiusstorage

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"

	"github.com/zvfkjytytw/marius/proto/mulus/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	maxMsgSize = 1024 * 1024 * 1024
	dataParts  = 6
)

type mus struct {
	sync.Mutex
	address string
	sarcina int32
}

type legionarius struct {
	conn   *grpc.ClientConn
	client api.MulusAPIClient
}

func (s *Storage) AddMus(ctx context.Context, address string) error {
	if _, ok := s.castra[address]; ok {
		return fmt.Errorf("mus %s is already in use", address)
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallSendMsgSize(maxMsgSize),
			grpc.MaxCallRecvMsgSize(maxMsgSize),
		),
	}

	conn, err := grpc.NewClient(address, opts...)

	if err != nil {
		return fmt.Errorf("fail dial to %s: %w", address, err)
	}

	s.Lock()
	defer s.Unlock()

	s.castra[address] = &legionarius{
		conn:   conn,
		client: api.NewMulusAPIClient(conn),
	}

	s.cohors = append(
		s.cohors,
		&mus{
			address: address,
			sarcina: 0,
		},
	)

	return nil
}

func (s *Storage) selectMulus() ([dataParts]*mus, error) {
	if len(s.cohors) < dataParts {
		return [dataParts]*mus{}, errors.New("not enough mulus in cohors")
	}

	var contubernium [dataParts]*mus
	s.RLock()
	copy(contubernium[:], s.cohors[:dataParts])
	s.RUnlock()

	for i := range contubernium {
		j := rand.Intn(i + 1)
		contubernium[i], contubernium[j] = contubernium[j], contubernium[i]
	}

	return contubernium, nil
}
