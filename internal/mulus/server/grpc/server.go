package mulusgrpc

import (
	"context"
	"fmt"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/zvfkjytytw/marius/proto/mulus/api/v1"
)

const (
	maxMsgSize = 1024 * 1024 * 1024
)

type storage interface {
	DeleteData(string) error
	GetData(string) ([]byte, error)
	SaveData(string, []byte) error
}

type ConfigGRPC struct {
	Host string `yaml:"host"`
	Port int32  `yaml:"port"`
}

type ServerGRPC struct {
	api.UnimplementedMulusAPIServer
	listener net.Listener
	server   *grpc.Server
	logger   *zap.Logger
	storage  storage
}

func NewGRPCServer(config *ConfigGRPC, logger *zap.Logger, storage storage) (*ServerGRPC, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.Host, config.Port))
	if err != nil {
		logger.Sugar().Fatalf("failed to listen: %v", err)
	}
	opts := []grpc.ServerOption{
		grpc.MaxRecvMsgSize(maxMsgSize),
		grpc.MaxSendMsgSize(maxMsgSize),
	}
	server := grpc.NewServer(opts...)

	return &ServerGRPC{
		listener: listener,
		server:   server,
		logger:   logger,
		storage:  storage,
	}, nil
}

// Run - start the server
func (s *ServerGRPC) Run(ctx context.Context) error {
	api.RegisterMulusAPIServer(s.server, s)

	return s.server.Serve(s.listener)
}

// Stop - gracefully stop the server.
func (s *ServerGRPC) Stop(ctx context.Context) error {
	s.logger.Info("gracefully stopping mulus grpc server at address")
	s.server.GracefulStop()
	s.listener.Close()

	return nil
}

// Kill - force stop the server.
func (s *ServerGRPC) Kill() {
	s.logger.Info("forced shutdown mulus grpc server at address")
	s.server.Stop()
	s.listener.Close()
}
