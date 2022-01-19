package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/abhchand/libmatch/internal/commands"
	"github.com/abhchand/libmatch/internal/config"
	"github.com/abhchand/libmatch/pkg/version"
	"github.com/urfave/cli/v2"
)

func main() {

	app := cli.NewApp()
	app.Name = "libmatch"
	app.HelpName = filepath.Base(os.Args[0])
	app.Usage = "A library for solving matching problems"
	app.Description = "For documentation, visit https://github.com/abhchand/libmatch#README"
	app.Version = version.Version()
	app.EnableBashCompletion = true
	app.Flags = config.GlobalFlags

	app.Commands = []*cli.Command{
		&commands.SolveCommand,
		&commands.LsCommand,
	}

	// Customize the output of `-v` / `--version`
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Fprintf(c.App.Writer, "%v\n", version.Formatted())
	}

	// Customize Application help-text
	cli.AppHelpTemplate = `
Usage: {{.HelpName}} [GLOBAL OPTIONS] COMMAND [OPTIONS]

{{.Usage}}
https://github.com/abhchand/libmatch#README

{{if .Commands}}
COMMANDS:
{{range .Commands}}{{if not .HideHelp}}   {{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}{{if .VisibleFlags}}
GLOBAL OPTIONS:
   {{range .VisibleFlags}}{{.}}
   {{end}}{{end}}

Run 'libmatch COMMAND --help' for more information on a command.
	`

	// Customize command help-text
	cli.CommandHelpTemplate = `Usage: {{.HelpName}} [GLOBAL OPTIONS] COMMAND [OPTIONS]

{{.Usage}}
{{if .VisibleFlags}}
OPTIONS:
   {{range .VisibleFlags}}{{.}}
   {{end}}{{end}}
`

	// Start the CLI
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}

}
