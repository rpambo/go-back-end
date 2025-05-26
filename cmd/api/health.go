package main

import "net/http"

func (a *application) healthandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok!!"))
}