package main

import (
	"log"

	"github.com/rpambo/go-back-end/internal/env"
)

func main(){
	cnf := config{
		addr: env.GetString("ADDR", ":8080"),
	}
	
	app := application{
		config: cnf,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}