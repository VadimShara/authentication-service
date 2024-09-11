package main

import(
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var globalFlags = []cli.Flag{
	&cli.StringFlag{
		Name:    "log-level",
		Usage:   "log level",
		EnvVars: []string{"LOG_LEVEL"},
		Value:   "local",
	},
}

func main() {
	app := &cli.App{
		Usage: "auth-service API Service",
		Commands: []*cli.Command{
			&Cmd,
		},
		Flags: globalFlags,
		Before: cli.BeforeFunc(func(c *cli.Context) error {
			return logger.SetupLogger(c.String("log-level"))
		}),
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("error start service: %v", err)
	}
}