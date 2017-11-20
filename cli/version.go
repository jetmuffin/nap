package cmd

import (
	"github.com/urfave/cli"
	"github.com/JetMuffin/nap/pkg/version"
	"os"
)

func Version() cli.Command {
	return cli.Command{
		Name:	"version",
		Description: "show version",
		Usage: "display version info",
		Action: func(c *cli.Context) error {
			return version.FormatVersion(os.Stdout)
		},
	}
}
