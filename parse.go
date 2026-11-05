package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Definition struct {
	Timing [][]float64 `json:"timing"`
}

type Config struct {
	Definitions map[string]Definition `json:"definitions"`
}

func ParseJsonFile(path string) (*Config, error) {
	jsonFileBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read json file: %v", err)
	}
	var config Config
	err = json.Unmarshal(jsonFileBytes, &config)
	if err != nil {
		return nil, fmt.Errorf("parse json file: %v", err)
	}

	return &config, nil
}
