package models

import (
	"fmt"
	"net/http"
)

// Transaction is a struct to the transaction model
type Transaction struct {
	ID        int     `json:"id"`
	Title     string  `json:"title"`
	Amount    float32 `json:"amount"`
	CreatedAt string  `json:"created_at"`
}

// Transactions is a array of Transaction struct
type TransactionList struct {
	Transactions []Transaction `json:"transactions"`
}

func (t *Transaction) Bind(r *http.Request) error {
	if t.Title == "" {
		return fmt.Errorf("name is a required field")
	}

	return nil
}

func (*TransactionList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*Transaction) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
