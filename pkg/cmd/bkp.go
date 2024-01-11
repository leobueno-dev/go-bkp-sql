package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var BkpCommand = &cli.Command{
	Name:  "bkp",
	Usage: "Choose a database type to backup",
	Subcommands: []*cli.Command{
		{
			Name:  "mssql",
			Usage: "Backup a mssql server database ",
			Action: func(ctx *cli.Context) error {
				fmt.Println("Backup a mssql server database")
				return nil
			},
		},
		{
			Name:  "postgress",
			Usage: "Backup a postgress database",
			Action: func(ctx *cli.Context) error {
				fmt.Println("Backup a postgress database")
				return nil
			},
		},
		{
			Name:  "mysql",
			Usage: "Backup a mysql database",
			Action: func(ctx *cli.Context) error {
				fmt.Println("Backup a mysql database")
				return nil
			},
		},
	},
}
