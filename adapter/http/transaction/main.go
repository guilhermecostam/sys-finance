package transaction

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/guilhermecostam/sys-finance/model/transaction"
)

// transactions is a instance of transaction model
var transactions = transaction.Transactions{}

// GetTransactions is a method to get all transactions in the system
func GetTransactions(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		responseWriter.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	responseWriter.Header().Set("Content-type", "application/json")

	_ = json.NewEncoder(responseWriter).Encode(transactions)
}

// CreateTransactions is a method to create a transactions in the system
func CreateTransactions(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		responseWriter.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var body, _ = ioutil.ReadAll(request.Body)
	_ = json.Unmarshal(body, &transactions)

	fmt.Println(transactions)
}
