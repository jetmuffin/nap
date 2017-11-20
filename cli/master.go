package cli

import (
	"fmt"
	"github.com/JetMuffin/nap/pkg/config"
	"github.com/JetMuffin/nap/pkg/master"
	"github.com/urfave/cli"
)

func Master() cli.Command {
	cmd := cli.Command{
		Name:        "master",
		Usage:       "start an master node",
		Description: "run nap master command",
		Action:      StartMaster,
	}

	cmd.Flags = []cli.Flag{
		FlagListenAddr(),
		FlagMesosAddr(),
		FlagLogLevel(),
	}

	return cmd
}

func StartMaster(c *cli.Context) error {
	cfg, err := config.NewMasterConfig(c)
	if err != nil {
		return fmt.Errorf("Failed to parse config: %v", err)
	}

	setupLogger(cfg.LogLevel)

	masterNode, err := master.New(cfg)
	if err != nil {
		return fmt.Errorf("Error when initilize master node: %v", err)
	}

	return masterNode.Start()
}
