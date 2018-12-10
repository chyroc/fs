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
	app.Name = "fs server"
	app.Usage = "sync file/folder to client"
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "host", Usage: "which host to listen", Value: "", Destination: &host},
		cli.IntFlag{Name: "port", Usage: "which port to listen", Value: 1234, Destination: &port},
	}
	app.Action = func(c *cli.Context) error {
		fmt.Println("server start!")
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func main() {
	StartApp()
}
