package router

import (
	"lab7/handler"
	"net/http"
)

func InitRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/style.css", fs)
	http.Handle("/script.js", fs)

	http.HandleFunc("/api/init", handler.InitHandler)
	http.HandleFunc("/api/get", handler.GetHandler)
	http.HandleFunc("/api/post", handler.PostHandler)
}
