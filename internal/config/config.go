package config

import (
	"github.com/urfave/cli/v2"
)

type Config struct {
	Algo   string
	Debug  bool
	CliCtx *cli.Context
}

func NewConfig(ctx *cli.Context) *Config {
	c := &Config{
		Debug: ctx.Bool("debug"),
	}

	return c
}
