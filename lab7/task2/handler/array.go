package handler

import (
	"encoding/json"
	"lab7/logic"
	"net/http"
)

func InitHandler(w http.ResponseWriter, r *http.Request) {
	arr := logic.InitArray(10)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(arr)
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	result := logic.ProcessArray()

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(result)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	result := logic.ProcessArray()

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(result)
}
