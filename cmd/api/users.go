package main

import (
	"net/http"

	"github.com/rpambo/go-back-end/types"
)

type userKey string

var userCtx userKey = "user"

func getUserFromContext(r *http.Request) *types.User{
	user, _ := r.Context().Value(userCtx).(*types.User)
	return user
}