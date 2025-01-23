package main

import (
	"log"

	"github.com/Ayeye11/inv/internal/db"
	"github.com/Ayeye11/inv/internal/store"
	"github.com/joho/godotenv"
)

func main() {
	// env
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found")
	}

	// config
	cfg := config{
		addr:   getEnvString("ADDR", ""),
		dsn:    getEnvString("DSN", ""),
		jwtKey: getEnvString("TOKEN_KEY", ""),
	}

	// db
	psql := db.NewPgSQL(cfg.dsn)
	db, err := psql.Run()
	if err != nil {
		log.Fatal(err)
	}

	// migrations

	// storage
	storage := store.NewStorage(db, cfg.jwtKey)

	// app
	app := newApplication(cfg.addr, storage)
	log.Fatal(app.run())
}
