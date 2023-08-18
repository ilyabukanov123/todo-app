package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Server Server `json:"server"`
}

type Server struct {
	Host     string `json:"host"`
	PortHTTP int    `json:"port_http"`
}

func GetConfig(pathConfig string) (*Config, error) {

	config := &Config{}

	f, err := os.Open(pathConfig)
	if err != nil {
		return config, err
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(config)
	return config, err
}
