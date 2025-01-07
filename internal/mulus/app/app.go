package mulusapp

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"

	common "github.com/zvfkjytytw/marius/internal/common"
	mgrpc "github.com/zvfkjytytw/marius/internal/mulus/server/grpc"
	lstorage "github.com/zvfkjytytw/marius/internal/mulus/storage/local"
)

type contubernium interface {
	Run(context.Context) error
	Stop(context.Context) error
	Kill()
}

type Config struct {
	GRPCConfig    *mgrpc.ConfigGRPC `yaml:"grpc_config"`
	StorageConfig *lstorage.Config  `yaml:"storage_config"`
}

type App struct {
	logger   *zap.Logger
	centuria []contubernium
}

func NewApp(config Config) (*App, error) {
	// init logger
	logger, err := common.InitLogger("mulus")
	if err != nil {
		return nil, fmt.Errorf("failed init logger: %v", err)
	}

	// init storage
	storage, err := lstorage.NewStorage(*config.StorageConfig)
	if err != nil {
		return nil, fmt.Errorf("failed init logger: %v", err)
	}

	// init grpc server
	server, err := mgrpc.NewGRPCServer(config.GRPCConfig, logger, storage)
	if err != nil {
		return nil, fmt.Errorf("failed init http server: %v", err)
	}

	return &App{
		logger: logger,
		centuria: []contubernium{
			server,
		},
	}, nil
}

func NewAppFromConfig(configFile string) (*App, error) {
	confData, err := common.ReadConfigFile(configFile)
	if err != nil {
		return nil, err
	}

	config := Config{}
	err = yaml.Unmarshal(confData, &config)
	if err != nil {
		return nil, err
	}

	return NewApp(config)
}

func (a *App) Run(ctx context.Context) {
	defer a.logger.Sync()

	sigChanel := make(chan os.Signal, 1)
	signal.Notify(sigChanel,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	for _, signum := range a.centuria {
		go func(signum contubernium) {
			if err := signum.Run(ctx); err != nil {
				a.logger.Sugar().Errorf("some signum failed to start: %v", err)
			}
		}(signum)
	}

	stopSignal := <-sigChanel
	a.logger.Sugar().Debugf("Stop by %v", stopSignal)
	a.stopAll(ctx)
}

func (a *App) stopAll(ctx context.Context) {
	var wg sync.WaitGroup
	wg.Add(len(a.centuria))
	for _, signum := range a.centuria {
		go func(signum contubernium) {
			defer wg.Done()
			if err := signum.Stop(ctx); err != nil {
				signum.Kill()
			}
		}(signum)
	}

	wg.Wait()
}
