package main

import (
	"strings"
	"fmt"
	"os"
	"ifconfig-cli/formatter"
	"strconv"
	"github.com/urfave/cli"
	"ifconfig-cli/client"
)

const (
	errorExitCode = 1
	version = "0.0.0"
	appName = "main"
	shortDescription = "cli for requesting network interfaces information from specified server"
)

func Address(server string, port int) string {
	return "http://" + server + ":" + strconv.Itoa(port) + "/service"
}

func main() {
	var server string
	var port int

	app := cli.NewApp()
	app.Name = appName
	app.Version = version
	app.Usage = shortDescription

	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "server",
			Value: "localhost",
			Usage: "Requested server",
			Destination: &server,
		},
		cli.IntFlag{
			Name: "port",
			Value: 55555,
			Usage: "Server's port",
			Destination: &port,
		},
	}

	app.Commands = []cli.Command{
		{
			Name: "list",
			Usage: "list available network interfaces",
			UsageText: appName + " list",
			Action: func(c *cli.Context) error {
				addr := Address(server, port)
				list, err := client.ListInterfaces(addr)
				if (err != nil) {
					return cli.NewExitError(err.Error(), errorExitCode)
				}
				fmt.Println( strings.Join(list, ", ") )
				return nil
			},
		},
		{
			Name: "show",
			Usage: "show specified interface information",
			UsageText: appName + " show [interface name]",
			Action: func(c *cli.Context) error {
				addr := Address(server, port)
				if (c.NArg() < 1) {
					return cli.NewExitError("Empty argument list for show command", 2)
				}
				info, err := client.ShowInterface(addr, c.Args().Get(0))
				if (err != nil) {
					return cli.NewExitError(err.Error(), errorExitCode)
				}
				err = formatter.PrintInterfaceInfo(info)
				if (err != nil) {
					return cli.NewExitError(err.Error(), errorExitCode)
				}
				return nil
			},
		},
	}
	app.Run(os.Args)
}
