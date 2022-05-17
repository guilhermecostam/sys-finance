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

// Bind is a method to validate required fields as requested
func (t *Transaction) Bind(r *http.Request) error {
	if t.Title == "" {
		return fmt.Errorf("name is a required field")
	}

	return nil
}

// Bind is a method to render TransactionList
func (*TransactionList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Bind is a method to render Transaction
func (*Transaction) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
