package config

import (
	"github.com/urfave/cli"
	"net/url"
)

type MasterConfig struct {
	ListenAddr string   `json:"listen_addr"`
	MesosAddr  *url.URL `json:"mesos_addr"`
	LogLevel   string
}

func NewMasterConfig(c *cli.Context) (*MasterConfig, error) {
	cfg := &MasterConfig{
		ListenAddr: "0.0.0.0:5678",
	}

	var err error

	cfg.MesosAddr, err = url.Parse(c.String("mesos"))
	if err != nil {
		return cfg, err
	}

	if c.String("loglevel") != "" {
		cfg.LogLevel = c.String("loglevel")
	}

	return cfg, nil
}
