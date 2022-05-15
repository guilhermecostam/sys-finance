package models

import (
	"time"
)

// Transaction is a struct to the transaction model
type Transaction struct {
	Title     string    `json:"title"`
	Amount    float32   `json:"amount"`
	Type      int       `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

// Transactions is a array of Transaction struct
type Transactions []Transaction
