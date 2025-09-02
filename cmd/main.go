package main

import (
	"gorestapi/config"
	"gorestapi/internals/repositories"
    "gorestapi/internals/services"
    "gorestapi/internals/handlers"
	"net/http"

)

func main() {
    postRepo := repositories.NewPostRepository(config.DB)
    postService := services.NewPostService(postRepo)

	handlers.RegisterRoutes("/v1", postService)
    http.ListenAndServe(":8080", nil)

}