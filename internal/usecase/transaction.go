package usecase

import (
	"github.com/lfsmariz/rinha-2024-q1/internal/dto"
	"github.com/lfsmariz/rinha-2024-q1/internal/repository"
)

func AddTransaction(id int64, t string, v int64, d string) (*dto.TransactionResponse, error) {
	r, err := repository.AddTransaction(id, t, v, d)

	return r, err
}

func GetBalanceAndLastTransactions(id int64) *dto.BankStatement {
	cB, cL := asyncGetBalance(id), asyncGetLastTransaction(id)

	b, l := <-cB, <-cL

	r := dto.BankStatement{
		Balance:         b,
		LastTransaction: l,
	}

	return &r
}

func asyncGetBalance(id int64) chan dto.Balance {
	r := make(chan dto.Balance)

	go func() {
		defer close(r)

		r <- *repository.GetBalance(id)
	}()

	return r
}

func asyncGetLastTransaction(id int64) chan []dto.LastTransaction {
	r := make(chan []dto.LastTransaction)

	go func() {
		defer close(r)

		r <- *repository.GetLastTransactions(id)
	}()

	return r
}
