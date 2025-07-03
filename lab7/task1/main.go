package main

import (
	"fmt"
	"net/http"

	"lab7/router"
)

func main() {
	router.InitRoutes()
	fmt.Println("Сервер запущено на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
