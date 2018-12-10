package main

import (
	"log"
	"os"

	"github.com/urfave/cli"

	"github.com/Chyroc/fs/internal/action"
)

func StartApp() {
	var host string
	var port int
	var dir string

	app := cli.NewApp()
	app.Name = "fs client"
	app.Usage = "sync file/folder from server"
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "host", Usage: "server host", Value: "", Destination: &host},
		cli.IntFlag{Name: "port", Usage: "server port", Value: 1234, Destination: &port},
		cli.StringFlag{Name: "dir", Usage: "which dir to sync", Value: ".", Destination: &dir},
	}
	app.Action = func(c *cli.Context) error {
		//directDir, err := filesys.GetDirectPath(dir)
		//if err != nil {
		//	return err
		//}

		return action.StartClient(host, port, dir)
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func main() {
	StartApp()
}
