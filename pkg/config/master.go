package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type MasterConfig struct {
	ListenAddr string `toml:"listen_addr"`
	MesosAddr  Addr   `toml:"mesos_addr"`
	LogLevel   string `toml:"log_level"`
	OAuthAddr  string `toml:"oauth_addr"`
}

func ParseMasterConfig(configPath string) (MasterConfig, error) {
	var cfg MasterConfig

	if _, err := toml.DecodeFile(configPath, &cfg); err != nil {
		return cfg, fmt.Errorf("parse config file error: %v", err)
	}

	return cfg, nil
}
