package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
)
const usage = `mydocer is a simple container runtime implementation 
 				the purpose of this project is to learn how docker works and how to 
write a docker by ourselves `

func main(){
	app := cli.NewApp()
	app.Name = "mydocker"
	app.Usage = usage

	app.Commands = []cli.Command{
		initCmd,
		runCmd,
	}
	app.Before = func(context *cli.Context) error {
		log.SetFormatter(&log.JSONFormatter{})
		log.SetOutput(os.Stdout)
		return nil
	}
	if err := app.Run(os.Args); err != nil{
		log.Fatal(err)
	}
}