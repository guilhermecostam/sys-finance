package http

import (
	"log"
	"net/http"

	"github.com/guilhermecostam/sys-finance/adapter/http/actuator"
	"github.com/guilhermecostam/sys-finance/adapter/http/transaction"
)

// Init is a function for init all endpoints of the system
func Init() {
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateTransactions)
	http.HandleFunc("/health", actuator.Health)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
