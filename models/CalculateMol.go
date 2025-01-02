package models

type MolRequest struct {
	Mass      float64 `json:"mass"`       // Maddenin kütlesi
	MolarMass float64 `json:"molarMass"` // Mol kütlesi
}

type MolResponse struct {
	Mol float64 `json:"mol"` // Hesaplanan mol
}