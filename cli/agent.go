package cli

import (
	"github.com/spf13/cobra"
)

type agentFlags struct {
	debug      bool
	configPath string
}

// AgentCommand implements agent subcommand.
type AgentCommand struct {
	baseCommand

	flags agentFlags
}

// Init initialize agent commands
func (a *AgentCommand) Init(c *Cli) {
	a.cli = c
	a.cmd = &cobra.Command{
		Use:   "agent",
		Short: "Start nap agent daemon",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return a.runAgent()
		},
	}
}

// addFlags add flags for specific command.
func (a *AgentCommand) addFlags() {
	flagSet := a.cmd.Flags()
	flagSet.BoolVar(&a.flags.debug, "debug", false, "Output debug log to stderr.")
	flagSet.StringVarP(&a.flags.configPath, "config", "c", "/etc/nap/config.toml", "Specify the path of configuration file.")
}

func (a *AgentCommand) runAgent() error {
	return nil
}
