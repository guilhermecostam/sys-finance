package database

import (
	"database/sql"

	"github.com/guilhermecostam/sys-finance/models"
	_ "github.com/lib/pq"
)

// CreateTransaction is a method for create a transaction in database
func (db Database) CreateTransaction(transaction *models.Transaction) error {
	var id int
	var createdAt string

	query := `INSERT INTO transactions (title, amount) VALUES ($1, $2) RETURNING id, created_at`
	err := db.Conn.QueryRow(query, transaction.Title, transaction.Amount).Scan(&id, &createdAt)
	if err != nil {
		return err
	}

	transaction.ID = id
	transaction.CreatedAt = createdAt

	return nil
}

// GetAllTransactions is a method for get all transactions in database
func (db Database) GetAllTransactions() (*models.TransactionList, error) {
	list := &models.TransactionList{}

	rows, err := db.Conn.Query("SELECT * FROM transactions ORDER BY ID")
	if err != nil {
		return list, err
	}

	for rows.Next() {
		var transaction models.Transaction
		err := rows.Scan(&transaction.Title, &transaction.Amount, &transaction.CreatedAt)
		if err != nil {
			return list, err
		}
		list.Transactions = append(list.Transactions, transaction)
	}

	return list, nil
}

// GetTransactionById is a method for get one transaction by id
func (db Database) GetTransactionById(transactionId int) (models.Transaction, error) {
	transaction := models.Transaction{}

	query := `SELECT * FROM transactions WHERE id = $1;`
	row := db.Conn.QueryRow(query, transactionId)

	switch err := row.Scan(
		&transaction.ID,
		&transaction.Title,
		&transaction.Amount,
		&transaction.CreatedAt); err {
	case sql.ErrNoRows:
		return transaction, ErrNoMatch
	default:
		return transaction, err
	}
}

// DeleteTransaction is a method for delete a transaction in database
func (db Database) DeleteTransaction(transactionId int) error {
	query := `DELETE FROM transactions WHERE id = $1;`
	_, err := db.Conn.Exec(query, transactionId)

	switch err {
	case sql.ErrNoRows:
		return ErrNoMatch
	default:
		return err
	}
}

// UpdateTransaction is a method for update a transaction in database
func (db Database) UpdateTransaction(transactionId int, itemData models.Transaction) (models.Transaction, error) {
	item := models.Transaction{}

	query := `UPDATE items SET name=$1, description=$2 WHERE id=$3 RETURNING id, name, description, created_at;`
	err := db.Conn.QueryRow(query, itemData.Title, itemData.Amount, transactionId).Scan(&item.ID, &item.Title, &item.Amount, &item.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return item, ErrNoMatch
		}
		return item, err
	}

	return item, nil
}
