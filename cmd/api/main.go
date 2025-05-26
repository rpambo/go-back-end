package main

import "log"

func main(){
	cnf := config{
		addr: ":8080",
	}
	app := application{
		config: cnf,
	}

	log.Fatal(app.run())
}