package dto

type TransactionRequest struct {
	Value       int64  `json:"valor" validate:"required,gte=0"`
	Type        string `json:"tipo" validate:"required,eq=c|eq=d"`
	Description string `json:"descricao" validate:"required,max=10"`
}

type TransactionResponse struct {
	Limit   int64 `json:"limite"`
	Balance int64 `json:"saldo"`
}
