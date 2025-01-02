package handlers

import (
	"encoding/json"
	"net/http"
	"omerozyilmaz/chemical-engineer-problem-solver/models"
	"omerozyilmaz/chemical-engineer-problem-solver/services"
)


func CalculateEquilibrium(w http.ResponseWriter, r *http.Request) {
    var req models.EquilibriumRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    result := services.CalculateEquilibrium(req)
    resp := models.EquilibriumResponse{Result: result}

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}