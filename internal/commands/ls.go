package commands

import (
	"fmt"
	"strings"

	"github.com/urfave/cli/v2"
)

var SEQ_BOLD = "\033[1m"
var SEQ_RESET = "\033[0m"

var AlgorithmDescriptions = map[string]string{
	"SRP": `Stable Roommates Problem
Find a stable matching within an even-sized set.
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
	fmt.Println("\nALGORITHMS:")
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
	}
	return nil
}
