package app

import "github.com/urfave/cli"

// Generate will return the command line application ready to be executed
func Generate() *cli.App {
	application := cli.NewApp()
	application.Name = "Command Line Application - Abacaxi"
	application.Usage = "Fetch IPs and Internet Server Names"
	return application
}
