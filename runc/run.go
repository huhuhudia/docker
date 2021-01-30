package main

import (
	"github.com/huhuhudia/docker/runc/container"
	"github.com/huhuhudia/docker/runc/def"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
)

var runCmd = cli.Command{
	Name: "run",
	Usage:`create a container with namespace and cgroup limit mydocker run -ti [command]`,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:"ti",
			Usage: "enable tty",
		},
	},
	Action: func(context *cli.Context) error{
		if len(context.Args()) < 1{
			return def.MissingArgsErr
		}
		cmd := context.Args().Get(0)
		tty := context.Bool("ti")
		Run(tty,cmd)
		return nil
	},
}

var initCmd = cli.Command{
	Name: "init",
	Usage: "init container process run users process in container . Do not call it outside ",
	Action: func(ctx *cli.Context) error{
		if len(ctx.Args()) < 1{
			return def.MissingArgsErr
		}
		cmd := ctx.Args().Get(0)
		log.Infoln("init command :%v", cmd)
		err := container.RunContainerInitProcess(cmd, nil)
		return err
	},
}

func Run(tty bool, command string){
	parent := container.NewParentProcess(tty, command)
	if err := parent.Start(); err != nil{
		log.Errorln(err)
	}
	parent.Wait()
	os.Exit(-1)
}