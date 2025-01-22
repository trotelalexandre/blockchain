package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/trotelalexandre/proto/node"
)

func LoadConfig(filename string) (*node.NodeConfig, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not open config file: %w", err)
	}
	defer file.Close()

	var config node.NodeConfig
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, fmt.Errorf("could not decode config file: %w", err)
	}

	return &config, nil
}