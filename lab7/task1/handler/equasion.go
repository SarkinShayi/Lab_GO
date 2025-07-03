package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"lab7/solver"
)

func SolveGetHandler(w http.ResponseWriter, r *http.Request) {
	a, _ := strconv.ParseFloat(r.URL.Query().Get("a"), 64)
	b, _ := strconv.ParseFloat(r.URL.Query().Get("b"), 64)
	c, _ := strconv.ParseFloat(r.URL.Query().Get("c"), 64)

	roots, err := solver.SolveQuadratic(a, b, c)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(roots)
}

func SolvePostHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		A float64 `json:"a"`
		B float64 `json:"b"`
		C float64 `json:"c"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		http.Error(w, "Неможливо прочитати JSON", http.StatusBadRequest)
		return
	}

	roots, err := solver.SolveQuadratic(req.A, req.B, req.C)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(roots)
}
