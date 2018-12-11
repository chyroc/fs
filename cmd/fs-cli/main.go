package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli"

	"github.com/Chyroc/fs/internal/action"
)

func StartApp() {
	var host string
	var port int
	var dir string
	var hot bool
	var ignore cli.StringSlice

	app := cli.NewApp()
	app.Name = "fs client"
	app.Usage = "sync file/folder from server"
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "host", Usage: "server host", Value: "", Destination: &host},
		cli.IntFlag{Name: "port", Usage: "server port", Value: 1234, Destination: &port},
		cli.StringFlag{Name: "dir", Usage: "which dir to sync", Value: ".", Destination: &dir},
		cli.BoolFlag{Name: "hot", Usage: "hot sync file while changed", Destination: &hot},
		cli.StringSliceFlag{Name: "ignore", Usage: "which files to ignore", Value: &ignore},
	}
	app.Action = func(c *cli.Context) error {
		if strings.HasPrefix(dir, "/") {
			return fmt.Errorf("please use relative path")
		}

		pwd, err := os.Getwd()
		if err != nil {
			return err
		}

		return action.StartClient(host, port, pwd, dir, hot, []string(ignore))
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func main() {
	StartApp()
}
