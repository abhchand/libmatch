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

	// Expand path of each `file` flag
	expandedFiles, err := expandFilenames(cfg.CliContext.StringSlice("file"))
	if err != nil {
		return cfg, err
	}
	cfg.Filenames = expandedFiles

	return cfg, nil
}

/*
 * Expands all relative filenames into absolute paths
 */
func expandFilenames(files []string) ([]string, error) {
	for f := range files {
		absFilename, err := filepath.Abs(files[f])

		if err != nil {
			return files, err
		}

		files[f] = absFilename
	}

	return files, nil
}
