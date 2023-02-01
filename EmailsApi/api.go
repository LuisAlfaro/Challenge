package main

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	r.Get("/search", getSearch)

	http.ListenAndServe(":3030", r)
}

func getSearch(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("text")
	var from int = 0
	var size int = 20

	fromString := r.URL.Query().Get("from")
	if fromString != "" {
		fromInt, err := strconv.Atoi(fromString)
		_ = err
		if fromInt > 0 {
			from = fromInt
		}
	}

	sizeString := r.URL.Query().Get("size")
	if sizeString != "" {
		sizeInt, err := strconv.Atoi(sizeString)
		_ = err
		if sizeInt > 0 {
			size = sizeInt
		}
	}

	//datos, err := zinc.Search(text, 0, 20)
	datos, err := Search(text, from, size)
	if err != nil {
		http.Error(w, http.StatusText(422), 422)
		return
	}

	w.Write(datos)

}
