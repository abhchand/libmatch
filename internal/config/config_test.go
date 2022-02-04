package config

import (
	"flag"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestNewConfig(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		flagSet := flag.NewFlagSet("test", 0)
		flagSet.String("algorithm", "SRP", "doc")
		flagSet.String("format", "csv", "doc")
		flagSet.Bool("debug", false, "doc")
		flagSet.Var(cli.NewStringSlice("/tmp/test.json"), "file", "doc")

		app := cli.NewApp()
		ctx := cli.NewContext(app, flagSet, nil)
		cfg, err := NewConfig(ctx)

		assert.Nil(t, err)

		assert.IsType(t, new(Config), cfg)
		assert.Equal(t, "SRP", cfg.Algorithm)
		assert.Equal(t, false, cfg.Debug)
		assert.Equal(t, "csv", cfg.OutputFormat)
		assert.Equal(t, []string{"/tmp/test.json"}, cfg.Filenames)
	})

	t.Run("`algorithm` is case insensitive", func(t *testing.T) {
		flagSet := flag.NewFlagSet("test", 0)
		flagSet.String("algorithm", "sRp", "doc")

		app := cli.NewApp()
		ctx := cli.NewContext(app, flagSet, nil)
		cfg, err := NewConfig(ctx)

		assert.Nil(t, err)

		assert.IsType(t, new(Config), cfg)
		assert.Equal(t, "SRP", cfg.Algorithm)
	})
}

func TestExpandFilenames(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		flagSet := flag.NewFlagSet("test", 0)
		flagSet.Var(cli.NewStringSlice("./test.json"), "file", "doc")

		app := cli.NewApp()
		ctx := cli.NewContext(app, flagSet, nil)
		cfg, err := NewConfig(ctx)

		curDir, err := filepath.Abs(".")

		assert.Nil(t, err)

		assert.IsType(t, new(Config), cfg)
		assert.Equal(t, []string{curDir + "/test.json"}, cfg.Filenames)
	})

	t.Run("handles multiple files", func(t *testing.T) {
		flagSet := flag.NewFlagSet("test", 0)
		flagSet.Var(cli.NewStringSlice("./a.json", "./b.json"), "file", "doc")

		app := cli.NewApp()
		ctx := cli.NewContext(app, flagSet, nil)
		cfg, err := NewConfig(ctx)

		curDir, err := filepath.Abs(".")

		assert.Nil(t, err)

		assert.IsType(t, new(Config), cfg)
		assert.Equal(t, []string{curDir + "/a.json", curDir + "/b.json"}, cfg.Filenames)
	})
}
