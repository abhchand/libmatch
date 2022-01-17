package config

import (
	"flag"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestNewConfig(t *testing.T) {
	ctx := CliTestContext()

	c := NewConfig(ctx)

	assert.IsType(t, new(Config), c)
	assert.Equal(t, false, c.Debug)
}

func CliTestContext() *cli.Context {
	globalSet := flag.NewFlagSet("test", 0)
	globalSet.Bool("debug", false, "doc")

	app := cli.NewApp()

	return cli.NewContext(app, globalSet, nil)
}
