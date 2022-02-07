package commands

import (
	"errors"
	"fmt"

	"github.com/abhchand/libmatch"
	"github.com/abhchand/libmatch/internal/config"
	"github.com/abhchand/libmatch/pkg/core"
	"github.com/abhchand/libmatch/pkg/load"
	"github.com/urfave/cli/v2"
)

type matchingAlgorithm struct {
	fileCount int
}

var MATCHING_ALGORITHMS_CFG = map[string]matchingAlgorithm{
	"SMP": {
		fileCount: 2,
	},
	"SRP": {
		fileCount: 1,
	},
}
var OUTPUT_FORMATS = [2]string{"csv", "json"}

/*
 * The `cli.Command` return value is wrapped in a function so we return a new
 * instance of it every time. This avoids caching flags between tests
 */
func SolveCommand() *cli.Command {
	return &cli.Command{
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
			&cli.StringSliceFlag{
				Name:     "file",
				Usage:    "JSON-formatted file containing list of matching preferences",
				Required: true,
				Aliases:  []string{"f"},
			},
			&cli.StringFlag{
				Name:     "format",
				Usage:    "Output format to print results. Must be one of 'csv', 'json'",
				Required: false,
				Value:    "csv",
				Aliases:  []string{"o"},
			},
		},
	}
}

func solveAction(ctx *cli.Context) error {
	var result core.MatchResult

	cfg, err := config.NewConfig(ctx)
	if err != nil {
		return err
	}

	if err = validateConfig(*cfg); err != nil {
		return err
	}

	prefsSet, err := loadFiles(*cfg)
	if err != nil {
		return err
	}

	/*
	 * `algorithm` value was already validated above, so it is guranteed
	 * to be one of the `case`s below.
	 */
	switch cfg.Algorithm {
	case "SMP":
		result, err = libmatch.SolveSMP(prefsSet[0], prefsSet[1])
	case "SRP":
		result, err = libmatch.SolveSRP(prefsSet[0])
	}

	if err != nil {
		return err
	}

	result.Print(cfg.OutputFormat)

	return nil
}

func validateConfig(cfg config.Config) error {
	mac := MATCHING_ALGORITHMS_CFG[cfg.Algorithm]

	// Verify `algorithm` value is valid
	valid := false
	if mac.fileCount == 0 {
		return errors.New(fmt.Sprintf("Unknown `--algorithm` value: %v", cfg.Algorithm))
	}

	// Verify number of `--file` inputs
	if len(cfg.Filenames) != mac.fileCount {
		return errors.New(
			fmt.Sprintf("Expected --file to be specified exactly %v time(s)", mac.fileCount))
	}

	// Verify `format` value is valid
	valid = false
	for i := range OUTPUT_FORMATS {
		if cfg.OutputFormat == OUTPUT_FORMATS[i] {
			valid = true
			break
		}
	}

	if !(valid) {
		return errors.New(fmt.Sprintf("Unknown `--format` value: %v", cfg.OutputFormat))
	}

	return nil
}

func loadFiles(cfg config.Config) ([]*[]core.MatchPreference, error) {
	prefsSet := make([]*[]core.MatchPreference, len(cfg.Filenames))

	for i := range cfg.Filenames {
		prefs, err := load.LoadFromFile(cfg.Filenames[i])

		if err != nil {
			return prefsSet, err
		}

		prefsSet[i] = prefs
	}

	return prefsSet, nil
}
