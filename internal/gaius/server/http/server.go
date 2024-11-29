package gaiushttp

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"
)

// type Storage interface {
// 	GetFile(string) ([]byte, error)
// 	SaveFile(string, []byte) (string, error)
//  UpdateFile(string, []byte) error
//  DeleteFile(string) error
// }

type FS interface {
	FindFile(string) (int32, error)
	CreateFile(string) (int32, error)
	DeleteFile(string) error
}

type ConfigHTTP struct {
	Host         string `yaml:"host"`
	Port         int32  `yaml:"port"`
	ReadTimeout  int32  `yaml:"read_timeout"`
	WriteTimeout int32  `yaml:"write_timeout"`
	IdleTimeout  int32  `yaml:"idle_timeout"`
}

type ServerHTTP struct {
	server *http.Server
	logger *zap.Logger
	fs     FS
}

func NewHTTPServer(config *ConfigHTTP, logger *zap.Logger, fs FS) (*ServerHTTP, error) {
	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", config.Host, config.Port),
		ReadTimeout:  time.Duration(config.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(config.IdleTimeout) * time.Second,
	}

	return &ServerHTTP{
		server: server,
		logger: logger,
		fs: fs,
	}, nil
}

// Run - start the server
func (s *ServerHTTP) Run(ctx context.Context) error {
	s.server.Handler = s.newRouter()
	s.logger.Sugar().Fatalf(
		"failed run http server: %v",
		s.server.ListenAndServe(),
	)
	return nil
}

// Stop - gracefully stop the server.
func (s *ServerHTTP) Stop(ctx context.Context) error {
	s.logger.Sugar().Infof("gracefully stopping http server at address %s", s.server.Addr)
	return s.server.Shutdown(ctx)
}

// Kill - force stop the server.
func (s *ServerHTTP) Kill() {
	s.logger.Sugar().Infof("forcefully stopping http server at address %s", s.server.Addr)
	s.server.Close()
}
