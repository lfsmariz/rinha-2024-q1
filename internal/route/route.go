package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/mux"
	"github.com/lfsmariz/rinha-2024-q1/internal/handler"
)

func CreateRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/clientes/{id:[1-5]}/transacoes", handler.AddTransaction).Methods("POST")
	r.HandleFunc("/clientes/{id:[1-5]}/extrato", handler.RetrieveBalance).Methods("GET")

	return r
}

func CreateRoutesFiber() *fiber.App {
	app := fiber.New(fiber.Config{
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Test App v1.0.1",
	})

	app.Get("/clientes/:id<max(5)>/extrato", handler.RetrieveBalanceFiber)
	app.Post("/clientes/:id<max(5)>/transacoes", handler.AddTransactionFiber)

	return app
}
