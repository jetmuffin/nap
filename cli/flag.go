package cli

import "github.com/urfave/cli"

func FlagListenAddr() cli.Flag {
	return cli.StringFlag{
		Name:   "listen",
		Usage:  "http listen endpoint address",
		EnvVar: "NAP_LISTEN_ADDR",
		Value:  "0.0.0.0:5678",
	}
}

func FlagMesosAddr() cli.Flag {
	return cli.StringFlag{
		Name:   "mesos",
		Usage:  "mesos address. e.g. mesos://host:port for standalone mode, zk://host1:port1,host2:port2,.../path for cluster mode",
		EnvVar: "NAP_MESOS_ADDR",
	}
}

func FlagOAuthAddr() cli.Flag {
	return cli.StringFlag{
		Name:   "oauth",
		Usage:  "oauth server address.",
		EnvVar: "NAP_OAUTH_ADDR",
		//Value: "127.0.0.1:3000",
	}
}

func FlagLogLevel() cli.Flag {
	return cli.StringFlag{
		Name:   "loglevel,l",
		Usage:  "customize log level [debug|info|error]",
		EnvVar: "NAP_LOG_LEVEL",
		Value:  "info",
	}
}
