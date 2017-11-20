package cli

import "github.com/urfave/cli"

func Main() {
	app := cli.NewApp()
	app.Name = "nap"
	app.Usage = "Next application platform"

	app.Commands = []cli.Command{
		Master(),
		Agent(),
		Version(),
	}

	app.RunAndExitOnError()
}
