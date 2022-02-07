package config

/*
 * Defines the `Config` structure that holds the internal `libmatch`
 * configurations
 */

import (
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v2"
)

type Config struct {
	Algorithm    string
	Debug        bool
	Filenames    []string
	OutputFormat string
	CliContext   *cli.Context
}

/*
 * Returns a new `Config` structure
 */
func NewConfig(ctx *cli.Context) (*Config, error) {
	cfg := &Config{
		Algorithm:    strings.ToUpper(ctx.String("algorithm")),
		Debug:        ctx.Bool("debug"),
		OutputFormat: ctx.String("format"),
		CliContext:   ctx,
	}

	if err := expandFilenames(cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}

/*
 * Expands all relative filenames into absolute paths
 */
func expandFilenames(cfg *Config) error {
	files := cfg.CliContext.StringSlice("file")

	for f := range files {
		absFilename, err := filepath.Abs(files[f])

		if err != nil {
			return err
		}

		files[f] = absFilename
	}

	cfg.Filenames = files

	return nil
}
