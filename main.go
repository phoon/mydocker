package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

const usage = "mydocker by go 1.11"

func main() {
	app := cli.NewApp()
	app.Name = "mydocker"
	app.Author = "Peven"
	app.Version = "Code-3.1"
	app.Usage = usage
	/* The commands to execute */
	app.Commands = []cli.Command{
		initCommand,
		runCommand,
	}

	app.Before = func(context *cli.Context) error {
		log.SetFormatter(&log.JSONFormatter{})
		log.SetOutput(os.Stdout)
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
