package main

import (
	"lab7/router"
	"log"
	"net/http"
)

func main() {
	router.InitRoutes()
	log.Println("Сервер запущено на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
