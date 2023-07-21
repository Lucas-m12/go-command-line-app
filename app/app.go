package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command line app"
	app.Usage = "Search ip and server names on internet"
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search Ip address on internet",
			Flags:  flags,
			Action: getIps,
		},
		{
			Name:   "servers",
			Usage:  "Get server names on internet",
			Flags:  flags,
			Action: getServers,
		},
	}
	return app
}

func getIps(context *cli.Context) {
	host := context.String("host")
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func getServers(context *cli.Context) {
	host := context.String("host")
	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}
	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
