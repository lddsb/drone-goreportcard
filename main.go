package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

// Version `cli version`
var Version = "20190309"

func main() {
	app := cli.NewApp()
	app.Name = "Drone Plugin GoReportCard"
	app.Usage = "fresh go report card message"
	app.Copyright = "© 2019 Dee Luo"
	app.Authors = []cli.Author{
		{
			Name:  "Dee Luo",
			Email: "luodi0128@gmail.com",
		},
	}
	app.Action = run
	app.Version = Version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "config.url",
			Usage:  "repository http url",
			EnvVar: "DRONE_GIT_HTTP_URL",
		},
	}

	if err := app.Run(os.Args); nil != err {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	p := Plugin{
		URL: c.String("config.url"),
	}
	return p.Exec()
}
