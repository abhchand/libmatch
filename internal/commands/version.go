package commands

import (
	"fmt"

	"github.com/abhchand/libmatch/pkg/version"
	"github.com/urfave/cli/v2"
)

// VersionCommand registers the version cli command.
var VersionCommand = cli.Command{
	Name:   "version",
	Usage:  "Print the version",
	Action: versionAction,
}

func versionAction(ctx *cli.Context) error {
	fmt.Println(version.Formatted())

	return nil
}
