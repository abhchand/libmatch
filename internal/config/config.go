package config

import (
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v2"
)

type Config struct {
	Algorithm       string
	Debug           bool
	Filename        string
	originalContext *cli.Context
}

func NewConfig(ctx *cli.Context) (*Config, error) {
	c := &Config{
		Algorithm: strings.ToUpper(ctx.String("algorithm")),
		Debug:     ctx.Bool("debug"),
	}

	absFilename, err := filepath.Abs(ctx.String("file"))
	if err != nil {
		return c, err
	}
	c.Filename = absFilename

	return c, nil
}
