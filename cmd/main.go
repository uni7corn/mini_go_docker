package main

import (
	"mini_go_docker/pkg/command"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

const usage = `mini docker is a simple container runtime implementation.
               The purpose of this projects is to learn how docker works and how to write a docker by ourselves
               Enjoy it, just for fun.`

func main() {
	logger := log.New()
	logger.Out = os.Stdout

	logger.Println("hello mini docker")
	app := cli.NewApp()
	app.Name = "mini docker"
	app.Usage = usage

	app.Commands = cli.Commands{command.InitCmd, command.RunCmd}

	app.Before = func(context *cli.Context) error {
		log.SetFormatter(&log.JSONFormatter{})
		return nil
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}

}
