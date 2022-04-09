package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name:  "management",
		Usage: "Manage the application",
		Commands: []*cli.Command{
			{
				Name:      "migrate",
				Usage:     "Control the migration flow of the project",
				ArgsUsage: "\n - up - migrating all the way up; \n - drop - dropping the databse; \n - <int> - migrating to selected version.",
				Action:    applyMigrations,
			},
			{
				Name:      "runserver",
				Usage:     "Run the server",
				ArgsUsage: "\n - port (default 4000)",
				Action:    runserver,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
