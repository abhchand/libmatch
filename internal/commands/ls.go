package commands

/*
 * Main handler for the `ls` subcommand.
 *
 * This subcommand lists all available matching algorithms.
 */

import (
	"fmt"
	"strings"

	"github.com/urfave/cli/v2"
)

var SEQ_BOLD = "\033[1m"
var SEQ_RESET = "\033[0m"

var AlgorithmDescriptions = map[string]string{
	"SMP": `Stable Marriage Problem

Find a stable matching between two same-sized sets.
Implements the Gale-Shapley (1962) algorithm.
A stable solution is always guranteed, but it is non-deterministic
and potentially one of many.

https://en.wikipedia.org/wiki/Stable_marriage_problem.`,
	"SRP": `Stable Roommates Problem

Find a stable matching within an even-sized set.
A stable solution is not guranteed, but is always deterministic if
exists.
Implements Irving's (1985) algorithm.

https://en.wikipedia.org/wiki/Stable_roommates_problem.
`,
}

/*
 * The `cli.Command` return value is wrapped in a function so we return a new
 * instance of it every time. This avoids caching flags between tests
 */
func LsCommand() *cli.Command {
	return &cli.Command{
		Name:   "ls",
		Usage:  "List all matching algorithms",
		Action: lsAction,
	}
}

func lsAction(ctx *cli.Context) error {
	fmt.Println("\nlibmatch supports the following matching algorithms:")
	fmt.Println("")

	for algorithm, desc := range AlgorithmDescriptions {
		num := 0
		lines := strings.Split(desc, "\n")

		for l := range lines {
			if num == 0 {
				fmt.Printf("\t%v%v%v\t\t%v\n", SEQ_BOLD, algorithm, SEQ_RESET, lines[l])
			} else {
				fmt.Printf("\t\t\t%v\n", lines[l])
			}

			num++
		}

		fmt.Println("")
		fmt.Println("")
	}
	return nil
}
