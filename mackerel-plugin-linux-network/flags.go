package main

import (
	"github.com/codegangsta/cli"
)

var cliFlags = []cli.Flag{
	cliFlagTempfile,
	cliFlagIpv4,
	cliFlagIpv6,
}

var cliFlagTempfile = cli.StringFlag{
	Name:   "tempfile",
	Value:  "/tmp/mackerel-plugin-linux-network",
	Usage:  "set temporary file path.",
	EnvVar: "ENVVAR_TEMPFILE",
}

var cliFlagIpv4 = cli.BoolFlag{
	Name:  "4",
	Usage: "print IPv4 metrics",
}

var cliFlagIpv6 = cli.BoolFlag{
	Name:  "6",
	Usage: "print IPv6 metrics",
}
