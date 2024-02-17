package route

import (
	"github.com/gorilla/mux"
	"github.com/lfsmariz/rinha-2024-q1/internal/handler"
)

func CreateRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/clientes/{id:[1-5]}/transacoes", handler.AddTransaction).Methods("POST")
	r.HandleFunc("/clientes/{id:[1-5]}/extrato", handler.RetrieveBalance).Methods("GET")

	return r
}
