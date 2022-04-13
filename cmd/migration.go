package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

func applyMigrations(c *cli.Context) error {

	dotenvErr := godotenv.Load(".env")
	if dotenvErr != nil {
		return dotenvErr
	}

	var postgresUrl string
	if c.String("environment") == "prod" {
		fmt.Println("migrating prod")
		postgresUrl = fmt.Sprintf("%s?sslmode=disable", os.Getenv("DATABASE_URL"))
	} else if c.String("environment") == "test" {
		postgresUrl = fmt.Sprintf("%s?sslmode=disable", os.Getenv("DATABASE_TEST_URL"))
	} else {
		fmt.Println(c.String("environment"))
		panic("arguments: Incorrect environment value. `prod` and `test` available")
	}

	m, err := migrate.New("file://migrations", postgresUrl)
	if err != nil {
		return dotenvErr
	}

	direction := c.Args().Get(0)
	switch direction {

	case "up":
		err = m.Up()

	case "drop":
		err = m.Drop()

	default:
		num, err := strconv.Atoi(direction)
		if err != nil || num < 0 {
			return errors.New("Provided direction must be 'up', 'drop' or a positive integer")
		}
		err = m.Migrate(uint(num))
	}
	if err != nil {
		return err
	}

	return nil
}
