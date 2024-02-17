package repository

import (
	"database/sql"
	"log"
	"time"

	"github.com/lfsmariz/rinha-2024-q1/internal/dto"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func Connection() {
	connStr := "host=localhost port=5432 user=admin password=123 dbname=rinha sslmode=disable"
	conn, err := sql.Open("postgres", connStr)

	Db = conn

	if err != nil {
		log.Fatal(err)
	}
}

func AddTransaction(id int64, t string, v int64, d string) (*dto.TransactionResponse, error) {
	row := Db.QueryRow("SELECT * FROM add_transaction($1, $2, $3, $4)", id, t, v, d)

	r := dto.TransactionResponse{}

	err := row.Scan(&r.Limit, &r.Balance)
	if err != nil {
		return nil, err
	}

	return &r, err
}

func GetBalance(id int64) *dto.Balance {
	row := Db.QueryRow("select limite, saldo from clientes c where c.id = $1", id)

	r := dto.Balance{BalanceDate: time.Now()}

	row.Scan(&r.Limit, &r.Total)

	return &r
}

func GetLastTransactions(id int64) *[]dto.LastTransaction {
	q := `select valor, tipo , descricao , realizada_em from transacoes t 
	where t.cliente_id = $1
	order by  t.realizada_em desc
	limit 10;`
	rows, err := Db.Query(q, id)

	s := make([]dto.LastTransaction, 0, 10)

	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		v := dto.LastTransaction{}

		if err := rows.Scan(&v.Value, &v.Type, &v.Description, &v.Date); err != nil {
			log.Println(err)
		}

		s = append(s, v)
	}

	return &s
}
