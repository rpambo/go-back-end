package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type application struct {
	config config
}

type config struct {
	addr string
}

func (a *application) mount() http.Handler {
	r := chi.NewRouter()

	// A good base middleware stack
  	r.Use(middleware.RequestID)
  	r.Use(middleware.RealIP)
  	r.Use(middleware.Logger)
  	r.Use(middleware.Recoverer)

  	// Set a timeout value on the request context (ctx), that will signal
  	// through ctx.Done() that the request has timed out and further
  	// processing should be stopped.
  	r.Use(middleware.Timeout(60 * time.Second))
  	
  	// Basic CORS
  	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
  	r.Use(cors.Handler(cors.Options{
    	// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
    	AllowedOrigins:   []string{"https://*", "http://*"},
    	// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
    	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    	AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
    	ExposedHeaders:   []string{"Link"},
    	AllowCredentials: false,
    	MaxAge:           300, // Maximum value not ignored by any of major browsers
  	}))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", a.healthandler)
	})

	return r
}

func (a *application) run(mux http.Handler) error {
	srv := http.Server{
		Addr: a.config.addr,
		Handler: mux,
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 10,
		IdleTimeout: time.Minute,
	}

	fmt.Printf("server listen %v", a.config.addr)

	return srv.ListenAndServe()
}