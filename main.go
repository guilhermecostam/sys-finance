package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/guilhermecostam/sys-finance/model/transaction"
)

func main() {
	http.HandleFunc("/transactions", getTransactions)
	http.HandleFunc("/transactions/create", createTransactions)

	http.ListenAndServe(":8000", nil)
}

func getTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("content-type", "application/json")

	const layout = "2006-01-02T15:04:05"
	// the '_' is for ignoring any error returns from the function
	salaryreceived, _ := time.Parse(layout, "2022-05-12T15:33:00")

	var transactions = transaction.Transactions{
		transaction.Transaction{
			Title:     "Salary",
			Amount:    1200.0,
			Type:      0,
			CreatedAt: salaryreceived,
		},
	}

	_ = json.NewEncoder(w).Encode(transactions)
}

func createTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var res = transaction.Transactions{}
	var body, _ = ioutil.ReadAll(r.Body)

	_ = json.Unmarshal(body, &res)
}
