package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func StartApp() {
	var host string
	var port int

	app := cli.NewApp()
	app.Name = "fs client"
	app.Usage = "sync file/folder from server"
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "host", Usage: "server host", Value: "", Destination: &host},
		cli.IntFlag{Name: "port", Usage: "server port", Value: 1234, Destination: &port},
	}
	app.Action = func(c *cli.Context) error {
		fmt.Println("client start!")
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func main() {
	StartApp()
}
