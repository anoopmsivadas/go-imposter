package httputils

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"

	"gitlab.com/anoopmsivadas/go-imposter/internal/pkg/config"
	"gitlab.com/anoopmsivadas/go-imposter/internal/pkg/logger"
	"gitlab.com/anoopmsivadas/go-imposter/internal/pkg/utils"
)

func CreateHandler(endpointConfig config.Endpoint) (func(writer http.ResponseWriter, req *http.Request), error) {
	appLogger, err := logger.GetInstance()
	if err != nil {
		fmt.Println(err)
		panic(0)
	}
	return func(writer http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		appLogger.Printf("%s: received %s:%s request\n", ctx.Value(utils.ServerAddrKey), req.Method, req.RequestURI)
		response, present := endpointConfig.ResponseMap[req.Method]
		writer.Header().Set(utils.ContentType, endpointConfig.ContentType)
		if req.URL.Path != endpointConfig.URI {
			writer.WriteHeader(http.StatusNotFound)
		} else if req.Method != endpointConfig.Method && present {
			io.WriteString(writer, response)
		} else if req.Method != endpointConfig.Method && !present {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			// io.WriteString(writer, fmt.Sprintf("{\"status\" : \"%s\"}", http.StatusText(http.StatusMethodNotAllowed)))
		} else {
			writer.WriteHeader(http.StatusOK)
			io.WriteString(writer, endpointConfig.Response)
		}
	}, nil
}

func CreateServer(ctx context.Context, handlerMap *map[string]func(w http.ResponseWriter, r *http.Request), serverPort uint16) *http.Server {
	mux := http.NewServeMux()
	for key, val := range *handlerMap {
		mux.HandleFunc(key, val)
	}

	serverPortStr := fmt.Sprintf(":%d", serverPort)

	server := &http.Server{
		Addr:    serverPortStr,
		Handler: mux,
		BaseContext: func(listener net.Listener) context.Context {
			ctx = context.WithValue(ctx, utils.ServerAddrKey, listener.Addr().String())
			return ctx
		},
	}
	return server
}
