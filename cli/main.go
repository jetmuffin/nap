package cli

import (
	"fmt"
	"os"
)

func Main() {
	cli := NewCli()

	base := &baseCommand{cmd: cli.rootCmd, cli: cli}

	// Add all subcommands
	cli.AddCommand(base, &VersionCommand{})
	cli.AddCommand(base, &MasterCommand{})

	if err := cli.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
