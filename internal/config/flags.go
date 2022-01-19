package config

import (
	"github.com/urfave/cli/v2"
)

// Flags available globally, with all commands
var GlobalFlags = []cli.Flag{
	&cli.BoolFlag{
		Name:  "debug",
		Usage: "enable debug output to STDOUT",
		Value: false,
	},
}
