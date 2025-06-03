package mulusapp

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	// 4 launch in docker compose
	"bytes"
	"net/http"
	"time"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"

	common "github.com/zvfkjytytw/marius/internal/common"
	mgrpc "github.com/zvfkjytytw/marius/internal/mulus/server/grpc"
	lstorage "github.com/zvfkjytytw/marius/internal/mulus/storage/local"
)

// 4 launch in docker compose
const (
	envGaiusHost = "GAIUS_HOST"
	envMusHost   = "MUS_HOST"
	gaiusPort    = 9090
)

var musPort int32

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
	logger, err := common.InitLogger("mus")
	if err != nil {
		return nil, fmt.Errorf("failed init logger: %v", err)
	}

	// init storage
	storage, err := lstorage.NewStorage(*config.StorageConfig)
	if err != nil {
		return nil, fmt.Errorf("failed init storage: %v", err)
	}

	// init grpc server
	server, err := mgrpc.NewGRPCServer(config.GRPCConfig, logger, storage)
	if err != nil {
		return nil, fmt.Errorf("failed init grpc server: %v", err)
	}

	// 4 launch in docker compose
	musPort = config.GRPCConfig.Port

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

	// 4 launch in docker compose
	if err := iungereToGaius(); err != nil {
		a.logger.Sugar().Errorf("failed to register to gaius service: %v", err)
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

// 4 docker compose running
func iungereToGaius() error {
	gaiusHost := os.Getenv(envGaiusHost)
	if gaiusHost == "" {
		return fmt.Errorf("gaius host is absent")
	}

	musHost := os.Getenv(envMusHost)
	if musHost == "" {
		return fmt.Errorf("mus host is absent")
	}

	if gaiusPort == 0 || musPort == 0 {
		return fmt.Errorf("gaius or Mus port is absent")
	}

	// wait for gaius is up in docker compose
	time.Sleep(time.Second * 2)

	// sent registration request to gaius
	url := fmt.Sprintf("http://%s:%d/add_mus", gaiusHost, gaiusPort)
	jsonStr := []byte(fmt.Sprintf(`{"address": "%s:%d"}`, musHost, musPort))
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed registration: %v", err)
	}
	defer resp.Body.Close()

	return nil
}
