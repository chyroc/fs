package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"

	"github.com/Chyroc/fs/internal/action"
)

func StartApp() {
	var host string
	var port int
	var mode string

	app := cli.NewApp()
	app.Name = "fs server"
	app.Usage = "sync file/folder to client"
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "host", Usage: "which host to listen", Value: "", Destination: &host},
		cli.IntFlag{Name: "port", Usage: "which port to listen", Value: 1234, Destination: &port},
		cli.StringFlag{Name: "mode", Usage: "pull or push", Value: "pull", Destination: &mode},
	}
	app.Action = func(c *cli.Context) error {
		fmt.Println("start with mode:", mode)
		fmt.Printf("address: %s:%d\n", host, port)

		switch mode {
		case "push":
			return action.StartFolderSync(mode, port)
		case "pull":
			return action.StartFolderSync(mode, port)
		default:
			return fmt.Errorf("mode must be pull or push!")
		}
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func main() {
	StartApp()
}
