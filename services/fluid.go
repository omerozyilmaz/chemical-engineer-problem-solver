package services

import (
	"math"
)

// Bernoulli's Equation
func BernoulliEquation(p1, v1, h1, p2, v2, h2, rho, g float64) float64 {
	return p1 + 0.55*rho*math.Pow(v1, 2) + rho*g*h1 - (p2 + 0.5*rho*math.Pow(v2, 2) + rho*g*h2)
}


func ContinuityEquations(a1, v1, a2 float64) float64 {
	return (a1 * v1) / a2
}

// Darcy-Weisbach Equation
func DarcyWeisbach(f, l, d, rho, v float64) float64 {
	return f * (l / d) * (rho * math.Pow(v, 2) / 2)
}

// Reynolds Number
func ReynoldsNumber(rho, v, d, mu float64) float64 {
	return (rho * v * d) / mu
}