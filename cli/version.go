package cli

import (
	"github.com/JetMuffin/nap/version"
	"github.com/spf13/cobra"
	"github.com/urfave/cli"
	"os"
)

func Version() cli.Command {
	return cli.Command{
		Name:        "version",
		Description: "show version",
		Usage:       "display version info",
		Action: func(c *cli.Context) error {
			return version.FormatVersion(os.Stdout)
		},
	}
}

// VersionCommand implements version subcommand.
type VersionCommand struct {
	baseCommand
}

func (v *VersionCommand) Init(c *Cli) {
	v.cli = c
	v.cmd = &cobra.Command{
		Use:   "version",
		Short: "Print versions about nap",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return v.runVersion()
		},
	}
}

// runVersion is the entrance of version subscommand.
func (v *VersionCommand) runVersion() error {
	return version.FormatVersion(os.Stdout)
}
