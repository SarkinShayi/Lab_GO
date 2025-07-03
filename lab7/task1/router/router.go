package router

import (
	"net/http"

	"lab7/handler"
)

func InitRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/style.css", fs)
	http.Handle("/script.js", fs)

	http.HandleFunc("/solve/get", handler.SolveGetHandler)
	http.HandleFunc("/solve/post", handler.SolvePostHandler)
}
