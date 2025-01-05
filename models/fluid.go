package models

// Bernoulli's Equation Request and Response
type BernoulliRequest struct {
	P1  float64 `json:"p1"`  
	V1  float64 `json:"v1"`  
	H1  float64 `json:"h1"`  
	P2  float64 `json:"p2"`  
	V2  float64 `json:"v2"`  
	H2  float64 `json:"h2"`  
	Rho float64 `json:"rho"` 
	G   float64 `json:"g"`   
}

type BernoulliResponse struct {
	DeltaEnergy float64 `json:"deltaEnergy"` 
}

// Continuity Equation Request and Response
type ContinuityRequest struct {
	A1 float64 `json:"a1"` 
	V1 float64 `json:"v1"` 
	A2 float64 `json:"a2"` 
}

type ContinuityResponse struct {
	V2 float64 `json:"v2"` 
}

// Darcy-Weisbach Equation Request and Response
type DarcyWeisbachRequest struct {
	F    float64 `json:"f"`    
	L    float64 `json:"l"`    
	D    float64 `json:"d"`    
	Rho  float64 `json:"rho"`  
	V    float64 `json:"v"`    
}

type DarcyWeisbachResponse struct {
	DeltaP float64 `json:"deltaP"` 
}

// Reynolds Number Request and Response
type ReynoldsRequest struct {
	Rho float64 `json:"rho"` 
	V   float64 `json:"v"`   
	D   float64 `json:"d"`   
	Mu  float64 `json:"mu"`  
}

type ReynoldsResponse struct {
	Re float64 `json:"re"` 
}