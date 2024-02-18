package usecase

import (
	"errors"

	"github.com/lfsmariz/rinha-2024-q1/internal/dto"
	"github.com/lfsmariz/rinha-2024-q1/internal/repository"
)

func AddTransaction(id int64, t string, v int64, d string) (*dto.TransactionResponse, error) {
	r, err := repository.AddTransaction(id, t, v, d)

	return r, err
}

func GetBalanceAndLastTransactions(id int64) (*dto.BankStatement, error) {
	cB, eB := asyncGetBalance(id)
	cL, eL := asyncGetLastTransaction(id)

	b, l, eb, el := <-cB, <-cL, <-eB, <-eL

	if eb != nil || el != nil {
		return nil, errors.New("invalid balance")
	}

	r := dto.BankStatement{
		Balance:         b,
		LastTransaction: l,
	}

	return &r, nil
}

func asyncGetBalance(id int64) (chan dto.Balance, chan error) {
	r := make(chan dto.Balance)
	err := make(chan error)

	go func() {
		defer close(r)
		defer close(err)

		v, e := repository.GetBalance(id)

		r <- *v
		err <- e
	}()

	return r, err
}

func asyncGetLastTransaction(id int64) (chan []dto.LastTransaction, chan error) {
	r := make(chan []dto.LastTransaction)
	err := make(chan error)
	go func() {
		defer close(r)
		defer close(err)

		v, e := repository.GetLastTransactions(id)

		r <- *v
		err <- e
	}()

	return r, err
}
