package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/zvfkjytytw/marius/proto/mulus/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	maxMsgSize = 1024 * 1024 * 1024
)

func main() {
	var mulusServer, testFile string
	flag.StringVar(&mulusServer, "s", "localhost:13081", "Mulus server")
	flag.StringVar(&testFile, "f", "test.file", "Test file")
	flag.Parse()

	data, err := os.ReadFile(testFile)
	if err != nil {
		fmt.Printf("failed read file %s: %v", testFile, err)
		os.Exit(1)
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallSendMsgSize(maxMsgSize),
			grpc.MaxCallRecvMsgSize(maxMsgSize),
		),
	}

	conn, err := grpc.NewClient(mulusServer, opts...)

	if err != nil {
		fmt.Printf("fail to dial: %v", err)
	}

	defer conn.Close()

	client := api.NewMulusAPIClient(conn)

	_, err = client.SaveData(
		context.Background(),
		&api.SaveRequest{
			Name: "testFile",
			Data: data,
		},
	)

	if err != nil {
		fmt.Printf("fail to save data: %v", err)
		os.Exit(1)
	}

	response, err := client.GetData(
		context.Background(),
		&api.GetRequest{
			Name: "testFile",
		},
	)

	if err != nil {
		fmt.Printf("fail to get data: %v", err)
		os.Exit(1)
	}

	if err = os.WriteFile("response.file", response.Data, 0644); err != nil {
		fmt.Printf("failed write response data to file: %v", err)
		os.Exit(1)
	}
}
