package config

import (
	"flag"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestNewConfig(t *testing.T) {
	globalSet := flag.NewFlagSet("test", 0)
	globalSet.String("algorithm", "SRP", "doc")
	globalSet.Bool("debug", false, "doc")
	globalSet.String("file", "/tmp/test.json", "doc")

	app := cli.NewApp()
	ctx := cli.NewContext(app, globalSet, nil)
	c, err := NewConfig(ctx)

	assert.Nil(t, err)

	assert.IsType(t, new(Config), c)
	assert.Equal(t, "SRP", c.Algorithm)
	assert.Equal(t, false, c.Debug)
	assert.Equal(t, "/tmp/test.json", c.Filename)
}

func TestNewConfig__AlgorithmCaseInsensitivity(t *testing.T) {
	globalSet := flag.NewFlagSet("test", 0)
	globalSet.String("algorithm", "sRp", "doc")

	app := cli.NewApp()
	ctx := cli.NewContext(app, globalSet, nil)
	c, err := NewConfig(ctx)

	assert.Nil(t, err)

	assert.IsType(t, new(Config), c)
	assert.Equal(t, "SRP", c.Algorithm)
}

func TestNewConfig__PathExpansion(t *testing.T) {
	globalSet := flag.NewFlagSet("test", 0)
	globalSet.String("file", "./test.json", "doc")

	app := cli.NewApp()
	ctx := cli.NewContext(app, globalSet, nil)
	c, err := NewConfig(ctx)

	curDir, err := filepath.Abs(".")

	assert.Nil(t, err)

	assert.IsType(t, new(Config), c)
	assert.Equal(t, curDir+"/test.json", c.Filename)
}
