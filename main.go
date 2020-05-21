package main

import (
	"net/http"
	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.GET("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome!"))
	})
	http.ListenAndServe(":8000", r)
}