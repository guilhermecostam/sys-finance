package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/guilhermecostam/sys-finance/database"
	"github.com/guilhermecostam/sys-finance/models"
)

var transactionIDKey = "transactionID"

func transactions(router chi.Router) {
	router.Get("/", getAllTransactions)
	router.Post("/", createTransaction)
	router.Route("/{transactionId}", func(router chi.Router) {
		router.Use(TransactionContext)
		router.Get("/", getTransaction)
		router.Put("/", updateTransaction)
		router.Delete("/", deleteTransaction)
	})
}

func TransactionContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		transactionId := chi.URLParam(r, "transactionId")
		if transactionId == "" {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("transaction ID is required")))
			return
		}

		id, err := strconv.Atoi(transactionId)
		if err != nil {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("invalid transaction ID")))
		}

		ctx := context.WithValue(r.Context(), transactionIDKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func createTransaction(w http.ResponseWriter, r *http.Request) {
	transaction := &models.Transaction{}

	if err := render.Bind(r, transaction); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	if err := dbInstance.CreateTransaction(transaction); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}

	if err := render.Render(w, r, transaction); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func getAllTransactions(w http.ResponseWriter, r *http.Request) {
	transactions, err := dbInstance.GetAllTransactions()
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}

	if err := render.Render(w, r, transactions); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
}

func getTransaction(w http.ResponseWriter, r *http.Request) {
	transactionID := r.Context().Value(transactionIDKey).(int)

	transaction, err := dbInstance.GetTransactionById(transactionID)
	if err != nil {
		if err == database.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ErrorRenderer(err))
		}
		return
	}

	if err := render.Render(w, r, &transaction); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func deleteTransaction(w http.ResponseWriter, r *http.Request) {
	transactionId := r.Context().Value(transactionIDKey).(int)

	err := dbInstance.DeleteTransaction(transactionId)
	if err != nil {
		if err == database.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
}
func updateTransaction(w http.ResponseWriter, r *http.Request) {
	transactionId := r.Context().Value(transactionIDKey).(int)

	transactionData := models.Transaction{}
	if err := render.Bind(r, &transactionData); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	transaction, err := dbInstance.UpdateTransaction(transactionId, transactionData)
	if err != nil {
		if err == database.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}

	if err := render.Render(w, r, &transaction); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}
