package main

import (
	"net/http"

	"github.com/rpambo/go-back-end/types"
)

type PostKey string
var postCtx PostKey = "post" 

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


func (app *application) GetByIdHandler(w http.ResponseWriter, r *http.Request){
	posts := getPostFromContext(r)

	comments, err := app.store.Comments.GetPostByID(r.Context(), posts.ID)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	posts.Comments = comments

	if err := app.jsonResponse(w, http.StatusOK, posts); err != nil{
		app.internalServerError(w, r, err)
		return
	}
}

func getPostFromContext(r *http.Request) (*types.Post){
	post, _ := r.Context().Value(postCtx).(*types.Post)

	return post
}