package gaiusapp

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
	ghttp "github.com/zvfkjytytw/marius/internal/gaius/server/http"
	gstorage "github.com/zvfkjytytw/marius/internal/gaius/storage"
	pg "github.com/zvfkjytytw/marius/internal/pg"
)

type Manipula interface {
	Run(context.Context) error
	Stop(context.Context) error
	Kill()
}

type Config struct {
	HTTPConfig *ghttp.ConfigHTTP `yaml:"http_config"`
	PGConfig   *pg.ConfigPG      `yaml:"pg_config"`
}

type App struct {
	logger *zap.Logger
	legio  []Manipula
}

func NewApp(config Config) (*App, error) {
	// init logger
	logger, err := common.InitLogger("gaius")
	if err != nil {
		return nil, fmt.Errorf("failed init logger: %v", err)
	}

	// init postgres agent
	pgAgent, err := pg.NewPGAgent(*config.PGConfig)
	if err != nil {
		return nil, fmt.Errorf("failed init postgres agent: %v", err)
	}

	// init storage
	storage, err := gstorage.NewStorage(pgAgent, logger)
	if err != nil {
		return nil, fmt.Errorf("failed init storage: %v", err)
	}

	// init http server
	server, err := ghttp.NewHTTPServer(config.HTTPConfig, logger, pgAgent, storage)
	if err != nil {
		return nil, fmt.Errorf("failed init http server: %v", err)
	}

	return &App{
		logger: logger,
		legio: []Manipula{
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

	for _, signum := range a.legio {
		go func(signum Manipula) {
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
	wg.Add(len(a.legio))
	for _, signum := range a.legio {
		go func(signum Manipula) {
			defer wg.Done()
			if err := signum.Stop(ctx); err != nil {
				signum.Kill()
			}
		}(signum)
	}

	wg.Wait()
}
