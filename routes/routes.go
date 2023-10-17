package routes

import (
	"github.com/gorilla/mux"
	"blog/controllers"
	"blog/middleware"
)
func SetupRouter() *mux.Router {
    r := mux.NewRouter()
    pc := controllers.PostController{}

    // Public routes (no token required)
    r.HandleFunc("/login", controllers.Login).Methods("POST")



    // Authenticated routes (require JWT token)
    authRouter := r.PathPrefix("/auth").Subrouter()
    authRouter.Use(middleware.AuthMiddleware)
		authRouter.HandleFunc("/fetchposts", pc.ListPosts).Methods("GET")
		authRouter.HandleFunc("/getpost/{id:[0-9]+}", pc.GetSinglePost).Methods("GET")
    authRouter.HandleFunc("/posts", pc.CreatePost).Methods("POST")
    authRouter.HandleFunc("/posts/{id:[0-9]+}", pc.UpdatePost).Methods("PUT")
    authRouter.HandleFunc("/posts/{id:[0-9]+}", pc.DeletePost).Methods("DELETE")

    return r
}
//curl -X POST http://localhost:8080/login -d "username=jprakash&password=pssnoida123"
