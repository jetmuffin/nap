package cli

import (
	"fmt"
	"github.com/JetMuffin/nap/pkg/config"
	"github.com/JetMuffin/nap/pkg/master"
	"github.com/spf13/cobra"
)

type masterFlags struct {
	debug      bool
	configFile string
}

// MasterCommand implements master subcommand.
type MasterCommand struct {
	baseCommand

	flags masterFlags
}

// Init initialize master commands
func (m *MasterCommand) Init(c *Cli) {
	m.cli = c
	m.cmd = &cobra.Command{
		Use:   "master",
		Short: "Start nap master daemon",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return m.runMaster()
		},
	}

	m.addFlags()
}

// addFlags adds flags for specific command.
func (m *MasterCommand) addFlags() {
	flagSet := m.cmd.Flags()
	flagSet.BoolVar(&m.flags.debug, "debug", false, "Output debug log to stderr.")
	flagSet.StringVarP(&m.flags.configFile, "config", "c", "/etc/nap/config.toml", "Specify the path of configuration file.")
}

func (m *MasterCommand) runMaster() error {
	cfg, err := config.ParseMasterConfig(m.flags.configFile)
	if err != nil {
		return err
	}

	setupLogger(cfg.LogLevel)

	master, err := master.New(cfg)
	if err != nil {
		return fmt.Errorf("Error when initilize master node: %v", err)
	}

	return master.Start()
}
