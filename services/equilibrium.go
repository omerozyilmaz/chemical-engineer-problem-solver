package services

import (
	"math"
	"omerozyilmaz/chemical-engineer-problem-solver/models"
)

func CalculateEquilibrium(req models.EquilibriumRequest) float64 {
    reactantProduct := 1.0
    productProduct := 1.0

    for i, val := range req.Reactants {
        reactantProduct *= math.Pow(val, float64(req.ReactantCoefficients[i]))
    }

    for i, val := range req.Products {
        productProduct *= math.Pow(val, float64(req.ProductCoefficients[i]))
    }

    return productProduct / reactantProduct
}