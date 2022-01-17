package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// SolveCommand registers the solve cli command.
var SolveCommand = cli.Command{
	Name:   "solve",
	Usage:  "Run a matching algorithm",
	Action: solveAction,
}

func solveAction(ctx *cli.Context) error {
	fmt.Println("TBD")

	return nil
}
