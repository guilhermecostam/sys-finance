package main

import (
	"net/http"

	"github.com/guilhermecostam/sys-finance/adapter/http/transaction"
)

func main() {
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateTransactions)

	http.ListenAndServe(":8000", nil)
}
