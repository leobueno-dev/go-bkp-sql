package main

import (
	"os"

	"github.com/leobueno-dev/go-bkp-sql/pkg/cmd"
	"github.com/leobueno-dev/go-bkp-sql/support/logger"
	zero "github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "go-bkp",
		Usage: "Make a sql database backup",
		Commands: []*cli.Command{
			cmd.BkpCommand,
		},
		Before: func(ctx *cli.Context) error {
			logger.SetupLogger("")
			log := logger.Logger("go-bkp", nil)

			ctx.Context = log.WithContext(ctx.Context)

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		zero.Fatal().Err(err).Msg("Fail run application")
	}
}
