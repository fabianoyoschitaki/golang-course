package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate will return the command line application ready to be executed
func Generate() *cli.App {
	application := cli.NewApp()
	application.Name = "Command Line Application - Abacaxi"
	application.Usage = "Fetch IPs and Internet Server Names"

	// common flags for commands
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",           // parameter name
			Value: "devbook.com.br", // default value
		},
	}
	application.Commands = []cli.Command{
		// e.g. go run main.go ip --host www.amazon.com.br
		{
			Name:   "ip",
			Usage:  "Fetch IPs given a hostname",
			Flags:  flags,
			Action: fetchIps,
		},
		// e.g. go run main.go servers --host www.amazon.com.br
		{
			Name:   "servers",
			Usage:  "Fetch servers in the Internet",
			Flags:  flags,
			Action: fetchNameServers,
		},
	}
	return application
}

// Fetch IPs given a host
func fetchIps(c *cli.Context) {
	host := c.String("host")

	ips, error := net.LookupIP(host)
	if error != nil {
		log.Fatal(error)
	}

	for _, ip := range ips {
		fmt.Printf("IP for host %s: %s\n", host, ip)
	}
}

// Fetch name servers given a host
func fetchNameServers(c *cli.Context) {
	host := c.String("host")

	servers, error := net.LookupNS(host) // name server
	if error != nil {
		log.Fatal(error)
	}

	for _, serverName := range servers {
		fmt.Printf("Server Name for host %s: %s\n", host, serverName.Host)
	}
}
