package main

import (
	"net/http"

	"github.com/rpambo/go-back-end/types"
)

func (app *application) CreatPost(w http.ResponseWriter, r *http.Request) {
	var payload types.CreatePostPayload

	if err := readJSON(w, r, &payload); err != nil{
		app.badRequestResponse(w, r, err)
		return
	}

	if err:= Validate.Struct(payload); err != nil{
		app.badRequestResponse(w, r, err)
		return
	}

	user := getUserFromContext(r)

	posts := &types.Post{
		Content: payload.Content,
		Title: payload.Title,
		Tags: payload.Tags,
		UserID: user.ID,
	}

	ctx := r.Context()

	if err := app.store.Posts.Create(ctx, posts); err != nil{
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, payload); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}