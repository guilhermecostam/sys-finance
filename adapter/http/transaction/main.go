package transaction

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/guilhermecostam/sys-finance/model/transaction"
	"github.com/guilhermecostam/sys-finance/util"
)

// GetTransactions is a method to get all transactions in the system
func GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("content-type", "application/json")

	var transactions = transaction.Transactions{
		transaction.Transaction{
			Title:     "Salary",
			Amount:    1200.0,
			Type:      0,
			CreatedAt: util.StringToTime("2022-05-12T15:33:00"),
		},
	}

	_ = json.NewEncoder(w).Encode(transactions)
}

// CreateTransactions is a method to create a transactions in the system
func CreateTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var res = transaction.Transactions{}
	var body, _ = ioutil.ReadAll(r.Body)

	_ = json.Unmarshal(body, &res)
}
