package main

import (
	"log"

	"github.com/rpambo/go-back-end/internal/db"
	"github.com/rpambo/go-back-end/internal/env"
	"github.com/rpambo/go-back-end/internal/store"
)

func main(){
	cnf := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr: env.GetString("DB_ADDR", "postgres://admin:admin@localhost?social?ssmode=disable"),
			maxOpenConns: env.GetInt("MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("MAX_IDLE_CONNS", 30),
			maxIdleTimes: env.GetString("MAX_IDLE_TIMES", "15m"),
		},
	}

	db, err := db.New(cnf.db.addr, cnf.db.maxOpenConns, cnf.db.maxIdleConns, cnf.db.maxIdleTimes)

	if err != nil{
		log.Panic(err)
	}
	
	store := store.NewStorage(db)

	app := application{
		config: cnf,
		store: store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}