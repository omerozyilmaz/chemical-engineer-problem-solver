package handlers

import (
	"encoding/json"
	"net/http"
	"omerozyilmaz/chemical-engineer-problem-solver/models"
	"omerozyilmaz/chemical-engineer-problem-solver/services"
)

func CalculateMol(w http.ResponseWriter, r *http.Request) {
	var req models.MolRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}


	result := services.CalculateMol(req)


	resp := models.MolResponse{Mol: result}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}