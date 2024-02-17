package dto

import "time"

type BankStatement struct {
	Balance         Balance           `json:"saldo"`
	LastTransaction []LastTransaction `json:"ultimas_transacoes"`
}

type Balance struct {
	Total       int64     `json:"total"`
	BalanceDate time.Time `json:"data_extrato"`
	Limit       int64     `json:"limite"`
}

type LastTransaction struct {
	Value       int64     `json:"valor"`
	Type        string    `json:"tipo"`
	Description string    `json:"descricao"`
	Date        time.Time `json:"realizada_em"`
}
