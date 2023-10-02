package config

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"gitlab.com/anoopmsivadas/go-imposter/internal/pkg/logger"
	"gitlab.com/anoopmsivadas/go-imposter/internal/pkg/utils"
)

type ConfigError struct{}

const configErrorMsg string = "Invalid Config"

func (err *ConfigError) Error() string {
	return configErrorMsg
}

func ReadConfig(filepath string) (Config, error) {
	appLogger, err := logger.GetInstance()
	if err != nil {
		fmt.Println(err)
		panic(0)
	}
	var configTest Config
	content, err := os.ReadFile(filepath)
	if err != nil {
		appLogger.Println(err)
		return configTest, &ConfigError{}
	}

	err = json.Unmarshal(content, &configTest)
	if err != nil {
		appLogger.Println(err)
		return configTest, &ConfigError{}
	}

	valid, normalized := validateNNormalize(&configTest)
	if !valid {
		return configTest, &ConfigError{}
	}
	return *normalized, nil
}

func validateNNormalize(conf *Config) (bool, *Config) {
	normalized := Config{}
	for _, service := range conf.Services {
		service_n := service
		endpointMap := make(map[string]*Endpoint)
		for _, endpoint := range service.Endpoints {
			if endpoint.Method == "" {
				return false, nil
			}

			if (endpoint.Method != http.MethodGet) && (endpoint.Method != http.MethodPost) && (endpoint.Method != http.MethodPut) && (endpoint.Method != http.MethodDelete) {
				return false, nil
			}

			if len(strings.TrimSpace(endpoint.URI)) == 0 {
				return false, nil
			}

			if val, present := endpointMap[endpoint.URI]; present {
				val.ResponseMap[endpoint.Method] = endpoint.Response
				// enableMethod(val)
				//endpoint.IgnoreItem = true
			} else {
				if endpoint.ResponseMap == nil {
					endpoint.ResponseMap = make(map[string]string)
				}
				endpoint_n := endpoint

				endpoint_n.ResponseMap[endpoint_n.Method] = endpoint_n.Response
				endpointMap[endpoint_n.URI] = &endpoint_n
				// enableMethod(&endpoint_n)
				service_n.Endpoints = append(service_n.Endpoints, endpoint_n)
			}
		}
		normalized.Services = append(normalized.Services, service_n)
	}

	return true, &normalized
}

func EnableMethod(endpoint *Endpoint) {
	if endpoint.Method == http.MethodGet {
		endpoint.EnabledMethods |= utils.GET
	}
	if endpoint.Method == http.MethodPost {
		endpoint.EnabledMethods |= utils.POST
	}
	if endpoint.Method == http.MethodPut {
		endpoint.EnabledMethods |= utils.PUT
	}
	if endpoint.Method == http.MethodDelete {
		endpoint.EnabledMethods |= utils.DELETE
	}
}
