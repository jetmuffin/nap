package cli

import (
	"github.com/JetMuffin/nap/version"
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
