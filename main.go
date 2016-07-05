package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/varhzj/easywebdemo/controllers/blog"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", blog.GetArticle).Methods(http.MethodGet)
	r.HandleFunc("/article/{id}", blog.ReadArticle).Methods(http.MethodGet)
	// user this to handle static files(js,css...)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./")))
	server := http.Server{
		Addr:    ":9000",
		Handler: r,
	}
	log.Fatal(server.ListenAndServe())
}
