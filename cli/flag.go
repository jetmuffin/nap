package cmd

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

func FlagLogLevel() cli.Flag {
	return cli.StringFlag{
		Name:   "loglevel,l",
		Usage:  "customize log level [debug|info|error]",
		EnvVar: "NAP_LOG_LEVEL",
		Value:  "info",
	}
}