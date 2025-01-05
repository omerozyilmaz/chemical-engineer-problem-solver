package handlers

import (
	"encoding/json"
	"net/http"
	"omerozyilmaz/chemical-engineer-problem-solver/models"
	"omerozyilmaz/chemical-engineer-problem-solver/services"
)

// Bernoulli's Equation Handler
func CalculateBernoulli(w http.ResponseWriter, r *http.Request) {
	var req models.BernoulliRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call Bernoulli Equation Service
	result := services.BernoulliEquation(req.P1, req.V1, req.H1, req.P2, req.V2, req.H2, req.Rho, req.G)
	response := models.BernoulliResponse{DeltaEnergy: result}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Continuity Equation Handler
func CalculateContinuity(w http.ResponseWriter, r *http.Request) {
	var req models.ContinuityRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call Continuity Equation Service
	result := services.ContinuityEquation(req.A1, req.V1, req.A2)
	response := models.ContinuityResponse{V2: result}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Darcy-Weisbach Equation Handler
func CalculateDarcyWeisbach(w http.ResponseWriter, r *http.Request) {
	var req models.DarcyWeisbachRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call Darcy-Weisbach Equation Service
	result := services.DarcyWeisbach(req.F, req.L, req.D, req.Rho, req.V)
	response := models.DarcyWeisbachResponse{DeltaP: result}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Reynolds Number Handler
func CalculateReynolds(w http.ResponseWriter, r *http.Request) {
	var req models.ReynoldsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call Reynolds Number Service
	result := services.ReynoldsNumber(req.Rho, req.V, req.D, req.Mu)
	response := models.ReynoldsResponse{Re: result}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}