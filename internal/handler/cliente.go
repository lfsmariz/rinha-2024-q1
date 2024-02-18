package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
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
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	err = validate.Struct(transaction)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
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
	res, err := usecase.GetBalanceAndLastTransactions(i)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	v, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(v)
}

// Fiber
func AddTransactionFiber(fc *fiber.Ctx) error {
	id := fc.Params("id")
	transaction := new(dto.TransactionRequest)

	if err := fc.BodyParser(transaction); err != nil {
		return fc.SendStatus(http.StatusUnprocessableEntity)
	}

	if err := validate.Struct(transaction); err != nil {
		return fc.SendStatus(http.StatusUnprocessableEntity)
	}

	i, _ := strconv.ParseInt(id, 10, 64)

	res, err := usecase.AddTransaction(i, transaction.Type, transaction.Value, transaction.Description)

	if err != nil {
		return fc.SendStatus(http.StatusUnprocessableEntity)
	}

	return fc.Status(http.StatusOK).JSON(res)
}

func RetrieveBalanceFiber(fc *fiber.Ctx) error {
	id := fc.Params("id")
	i, _ := strconv.ParseInt(id, 10, 64)
	res, err := usecase.GetBalanceAndLastTransactions(i)
	if err != nil {
		return fc.SendStatus(http.StatusUnprocessableEntity)
	}
	return fc.Status(http.StatusOK).JSON(res)
}
