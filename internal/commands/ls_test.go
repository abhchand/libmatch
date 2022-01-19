package commands

import (
	"flag"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestLsAction(t *testing.T) {
	globalSet := flag.NewFlagSet("test", 0)
	globalSet.Bool("debug", false, "doc")

	app := cli.NewApp()
	ctx := cli.NewContext(app, globalSet, nil)
	err := lsAction(ctx)

	assert.Nil(t, err)
}
