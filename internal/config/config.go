package config

import(
	"github.com/urfave/cli"
)

type Config struct {
	Algo     string
	Debug    bool
	CliCtx   *cli.Context
}

func NewConfig(ctx *cli.Context) *Config {
	c := &Config{
		Debug: ctx.GlobalBool("debug"),
	}

	return c
}
