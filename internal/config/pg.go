package config

import (
	"database/sql"
	"log"
	"time"
)

func PgConnection() *sql.DB {
	pgurl := PostgresUrl()
	pg, err := sql.Open("postgres", pgurl)
	if err != nil {
		log.Fatal(err)
	}
	pg.SetMaxIdleConns(5)
	pg.SetMaxOpenConns(20)
	pg.SetConnMaxLifetime(60 * time.Minute)
	pg.SetConnMaxIdleTime(10 * time.Minute)

	if err := pg.Ping(); err != nil {
		log.Fatal(err)
	}

	return pg
}
