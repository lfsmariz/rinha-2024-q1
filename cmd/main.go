package main

import (
	"fmt"
	"net/http"

	"github.com/lfsmariz/rinha-2024-q1/internal/repository"
	"github.com/lfsmariz/rinha-2024-q1/internal/route"
)

func main() {
	p := ":5000"
	r := route.CreateRoutes()
	repository.Connection()

	fmt.Println("Starting Server on port " + p)

	http.ListenAndServe(p, r)
}
