package handlers

import (
    "gorestapi/internals/handlers/v1"
    "gorestapi/internals/services"
    "net/http"
    "strings"
)

func RegisterRoutes(prefix string, postService *services.PostService) {
    // Ensure prefix starts with "/" and does not end with "/"
    if !strings.HasPrefix(prefix, "/") {
        prefix = "/" + prefix
    }
    prefix = strings.TrimSuffix(prefix, "/")

    postHandler := v1.NewPostHandler(postService)

    http.HandleFunc(prefix+"/posts/create", postHandler.Create)
    http.HandleFunc(prefix+"/posts/get", postHandler.GetByID)
    http.HandleFunc(prefix+"/posts/all", postHandler.GetAll)
    http.HandleFunc(prefix+"/posts/update", postHandler.Update)
    http.HandleFunc(prefix+"/posts/delete", postHandler.Delete)
}