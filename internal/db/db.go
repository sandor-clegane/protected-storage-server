package db

import (
	"database/sql"
	"embed"
	"fmt"
	"log"

	"github.com/pressly/goose/v3"
)

var db *sql.DB

//go:embed migrations/*.sql
var embedMigrations embed.FS

func InitDB(dbAddress string) (*sql.DB, error) {
	if db != nil {
		return db, nil
	}

	fmt.Println(dbAddress)
	newConn, connectionErr := sql.Open("postgres", dbAddress)
	if connectionErr != nil {
		log.Println(connectionErr)
		return nil, connectionErr
	}
	db = newConn

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}

	return db, nil
}
