package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	mApp "github.com/zvfkjytytw/marius/internal/mulus/app"
)

const envConfigFile = "APP_CONFIG_FILE"

func main() {
	var configFile string
	flag.StringVar(&configFile, "c", "./mulus.yaml", "Mulus config file")
	flag.Parse()

	if envCF := os.Getenv(envConfigFile); envCF != "" {
		configFile = envCF
	}

	app, err := mApp.NewAppFromConfig(configFile)
	if err != nil {
		fmt.Printf("failed create mulus app: %v\n", err)
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app.Run(ctx)
}
