package services

import "omerozyilmaz/chemical-engineer-problem-solver/models"


func CalculateMol(req models.MolRequest) float64 {

	return req.Mass / req.MolarMass
}