package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
)

type variables struct {
	Addr string
	DSN  string
}

func ParseVariables() variables {

	dotenvErr := godotenv.Load(".env")

	if dotenvErr != nil {
		log.Print("Error loading .env file")
	}
	addrInput := os.Getenv("ADDR")
	addr := fmt.Sprintf(":%v", addrInput)
	dsn := os.Getenv("DATABASE_URL")

	if dsn == "" {
		log.Fatal("No database url provided")
	}

	return variables{addr, dsn}
}

func OpenDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
