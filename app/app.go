package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate will return the object that will be our CLI to be executed
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Demo Command Line Interface"
	app.Usage = "Search for IP's and Names from intenet servers"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "opensource.janestreet.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search for IPs from addresses in the internet",
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:   "server",
			Usage:  "Search for Servers names from addresses in the internet",
			Flags:  flags,
			Action: searchServers,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	// Context -> cli.Command.cliFlag
	// Context.String -> StringFlag
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	// _ is to ignore the index
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServers(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server)
	}

}
