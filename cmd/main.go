package main

import (
	"github.com/lfsmariz/rinha-2024-q1/internal/repository"
	"github.com/lfsmariz/rinha-2024-q1/internal/route"
)

func main() {
	p := ":5003"
	repository.Connection()

	rf := route.CreateRoutesFiber()

	rf.Listen(p)
}
