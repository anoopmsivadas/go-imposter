package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"gitlab.com/anoopmsivadas/go-imposter/internal/pkg/config"
	"gitlab.com/anoopmsivadas/go-imposter/internal/pkg/httputils"
	"gitlab.com/anoopmsivadas/go-imposter/internal/pkg/logger"
)

func main() {

	logger.InitLogger(true, "./logs", logger.INFO)

	appLogger, err := logger.GetInstance()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	/*
		flag.Usage = func() {
			fmt.Fprint(os.Stderr, fmt.Sprint(cli.HelpText))
		}
		flag.Usage()
		os.Exit(0)

	*/
	serverList := make([]*http.Server, 0)
	ctx, cancelCtx := context.WithCancel(context.Background())

	appConfig, err := config.ReadConfig("config.json")
	if err != nil {
		appLogger.Printf("%s\n", err)
		cancelCtx()
		<-ctx.Done()
	}

	for _, service := range appConfig.Services {
		handlerMap := make(map[string]func(w http.ResponseWriter, r *http.Request))
		for _, endpoint := range service.Endpoints {
			handlerFun, err := httputils.CreateHandler(endpoint)
			if err == nil {
				handlerMap[endpoint.URI] = handlerFun
			}
		}
		_server := httputils.CreateServer(ctx, &handlerMap, service.Port)
		serverList = append(serverList, _server)
	}

	// var wg sync.WaitGroup
	for _, server := range serverList {
		go func() {
			// For loops will be fixed in Go v1.22, with 1.21 run with `GOEXPERIMENT=loopvar go test` env var set.
			// https://go.dev/blog/loopvar-preview
			appLogger.Println("spawning server on port:", server.Addr)
			err := server.ListenAndServe()
			if errors.Is(err, http.ErrServerClosed) {
				appLogger.Printf("server %s closed\n", server.Addr)
			} else if err != nil {
				appLogger.Printf("error listening for server: %s\n", err)
			}
			cancelCtx()
		}()
	}
	// wg.Wait()
	<-ctx.Done()

}
