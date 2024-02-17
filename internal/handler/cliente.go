package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/lfsmariz/rinha-2024-q1/internal/dto"
	"github.com/lfsmariz/rinha-2024-q1/internal/usecase"
)

var validate = validator.New()

func AddTransaction(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	decoder := json.NewDecoder(r.Body)
	var transaction dto.TransactionRequest

	err := decoder.Decode(&transaction)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = validate.Struct(transaction)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	i, _ := strconv.ParseInt(params["id"], 10, 64)

	res, err := usecase.AddTransaction(i, transaction.Type, transaction.Value, transaction.Description)

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	v, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(v)
}

func RetrieveBalance(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	i, _ := strconv.ParseInt(params["id"], 10, 64)
	res := usecase.GetBalanceAndLastTransactions(i)
	v, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(v)
}
