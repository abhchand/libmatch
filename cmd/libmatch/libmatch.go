package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/abhchand/libmatch/internal/commands"
	"github.com/abhchand/libmatch/internal/config"
	"github.com/abhchand/libmatch/pkg/version"
	"github.com/urfave/cli"
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

	app.Commands = []cli.Command{
		commands.SolveCommand,
		commands.VersionCommand,
	}

	if err := app.Run(os.Args); err != nil {
		// log.Error(err)
		fmt.Println(err)
	}

}
