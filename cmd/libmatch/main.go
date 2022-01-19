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
	app.Usage = "A Go library for solving matching problems"
	app.Description = "For documentation, visit https://github.com/abhchand/libmatch"
	app.Version = version.Version()
	app.EnableBashCompletion = true
	app.Flags = config.GlobalFlags

	app.Commands = []*cli.Command{
		&commands.SolveCommand,
		&commands.VersionCommand,
	}

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

	cli.CommandHelpTemplate = `Usage: {{.HelpName}} [GLOBAL OPTIONS] COMMAND [OPTIONS]

{{.Usage}}
{{if .VisibleFlags}}
OPTIONS:
   {{range .VisibleFlags}}{{.}}
   {{end}}{{end}}
`

	if err := app.Run(os.Args); err != nil {
		// log.Error(err)
		fmt.Println(err)
	}

}
