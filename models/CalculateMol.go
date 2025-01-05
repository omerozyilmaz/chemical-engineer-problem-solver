package models

type MolRequest struct {
	Mass      float64 `json:"mass"`     
	MolarMass float64 `json:"molarMass"`
}

type MolResponse struct {
	Mol float64 `json:"mol"` 
}