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
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))

}