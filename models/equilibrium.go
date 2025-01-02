package models

type EquilibriumRequest struct {
    Reactants            []float64 `json:"reactants"`
    Products             []float64 `json:"products"`
    ReactantCoefficients []int     `json:"reactantCoefficients"`
    ProductCoefficients  []int     `json:"productCoefficients"`
}

type EquilibriumResponse struct {
    Result float64 `json:"result"`
}