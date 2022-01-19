package commands

import (
	"errors"
	"fmt"

	"github.com/abhchand/libmatch"
	"github.com/abhchand/libmatch/internal/config"
	"github.com/abhchand/libmatch/pkg/load"
	"github.com/urfave/cli/v2"
)

var ALGORITHMS = [1]string{"SRP"}

// SolveCommand registers the solve cli command.
var SolveCommand = cli.Command{
	Name:   "solve",
	Usage:  "Run a matching algorithm",
	Action: solveAction,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "algorithm",
			Usage:    "Algorithm used to determine matches. See all algorithms with \"libmatch ls\"",
			Required: true,
			Aliases:  []string{"a"},
		},
		&cli.StringFlag{
			Name:     "file",
			Usage:    "JSON-formatted file containing list of matching preferences",
			Required: true,
			Aliases:  []string{"f"},
		},
	},
}

func solveAction(ctx *cli.Context) error {
	cfg, err := config.NewConfig(ctx)
	if err != nil {
		return err
	}

	if err = validateConfig(*cfg); err != nil {
		return err
	}

	prefs, err := load.LoadFromFile(cfg.Filename)
	if err != nil {
		return err
	}

	/*
	 * `algorithm` value was already validated above, so it is guranteed
	 * to be one of the `case`s below.
	 */
	switch cfg.Algorithm {
	case "SRP":
		_, err = libmatch.SolveSRP(prefs)
		if err != nil {
			return err
		}
	}

	return nil
}

func validateConfig(cfg config.Config) error {
	// Verify `algorithm` value is valid
	for i := range ALGORITHMS {
		if cfg.Algorithm == ALGORITHMS[i] {
			return nil
		}
	}

	return errors.New(fmt.Sprintf("Unknown `--algorithm` value: %v", cfg.Algorithm))
}
