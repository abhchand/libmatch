// Package config provides accessors to create and model internal libmatch
// configuration
package config

import (
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v2"
)

// Config defines the structure of the internal libmatch configuration
type Config struct {
	Algorithm    string
	Debug        bool
	Filenames    []string
	OutputFormat string
	CliContext   *cli.Context
}

// NewConfig returns a new Config structure
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

// expandFilenames expands all relative filenames into absolute paths
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
