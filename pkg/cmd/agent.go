package cmd

import "github.com/urfave/cli"

func Agent() cli.Command {
	agentCmd := cli.Command{
		Name:        "agent",
		Usage:       "start an agent node",
		Description: "run nap agent command",
		Action:      StartAgent,
	}

	agentCmd.Flags = []cli.Flag{
		FlagListenAddr(),
	}

	return agentCmd
}

func StartAgent(c *cli.Context) error {
	return nil
}
