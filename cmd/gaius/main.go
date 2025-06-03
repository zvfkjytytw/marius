package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	gApp "github.com/zvfkjytytw/marius/internal/gaius/app"
)

const envConfigFile = "APP_CONFIG_FILE"

func main() {
	var configFile string
	flag.StringVar(&configFile, "c", "./gaius.yaml", "Gaius config file")
	flag.Parse()

	if envCF := os.Getenv(envConfigFile); envCF != "" {
		configFile = envCF
	}

	app, err := gApp.NewAppFromConfig(configFile)
	if err != nil {
		fmt.Printf("failed create Gaius app: %v\n", err)
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app.Run(ctx)
}
