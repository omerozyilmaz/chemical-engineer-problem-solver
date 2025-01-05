package main

import (
	"fmt"
	"log"
	"net/http"
	"omerozyilmaz/chemical-engineer-problem-solver/handlers"
)

func main() {
    fmt.Println("Chemical Engineer Problem Solver API is running!")
	http.HandleFunc("/api/mol", handlers.CalculateMol)
	http.HandleFunc("/api/equilibrium", handlers.CalculateEquilibrium)
	http.HandleFunc("/api/bernoulli", handlers.CalculateBernoulli)
	http.HandleFunc("/api/continuity", handlers.CalculateContinuity)
	http.HandleFunc("/api/darcy", handlers.CalculateDarcyWeisbach)
	http.HandleFunc("/api/reynolds", handlers.CalculateReynolds)

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))

}