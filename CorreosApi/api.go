package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var config Config = Config{}
var zinc Zinc = Zinc{}

func main() {
	config, err := NewConfig("config,json")
	if err != nil {
		log.Fatal(err)
	}
	zinc = Zinc{
		Server:   config.ZincServer,
		Index:    config.ZincIndex,
		User:     config.ZincUser,
		Password: config.ZincPassword,
	}
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	r.Get("/search", getSearch)

	http.ListenAndServe(":3000", r)
}

func getSearch(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("Text")
	datos, err := zinc.Search(text)
	if err != nil {
		http.Error(w, http.StatusText(422), 422)
		return
	}
	w.Write(datos)

}
