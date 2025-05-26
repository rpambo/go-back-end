package main

import (
	"fmt"
	"net/http"
)

type application struct {
	config config
}

type config struct {
	addr string
}

func (a *application) run() error {
	srv := http.Server{
		Addr: a.config.addr,
	}

	fmt.Printf("server listen %v", a.config.addr)

	return srv.ListenAndServe()
}