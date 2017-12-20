package cli

import "github.com/spf13/cobra"

// Cli is the entrance of nap, it manage all commands.
type Cli struct {
	rootCmd *cobra.Command
}

// NewCli creeates an instance of Cli.
func NewCli() *Cli {
	return &Cli{
		rootCmd: &cobra.Command{
			Use:   "nap",
			Short: "Next application platform.",
		},
	}
}

// Run executes the nap program.
func (c *Cli) Run() error {
	return c.rootCmd.Execute()
}

// AddCommand add a subcommand.
func (c *Cli) AddCommand(parent, child Command) {
	child.Init(c)

	parentCmd := parent.Cmd()
	childCmd := child.Cmd()

	// make command error not return command usage and error
	childCmd.SilenceUsage = true
	childCmd.SilenceErrors = true

	childCmd.PreRun = func(cmd *cobra.Command, args []string) {}

	parentCmd.AddCommand(childCmd)
}
