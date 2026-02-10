package repositories

import (
	"categories-api/model"
	"database/sql"
	"fmt"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(dn *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: dn}
}

func (repo *TransactionRepository) CreateTransaction(items []model.CheckoutItem) (*model.Transaction, error) {
	if len(items) <= 0 {
		return nil, fmt.Errorf("Len items is zero")
	}
	var (
		res *model.Transaction
	)

	tx, err := repo.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	totalAmount := 0
	details := make([]model.TransactionDetail, 0)
	// loop setiap item
	for _, item := range items {
		var productName string
		var productID, price, stock int
		err := tx.QueryRow("SELECT id, name, price, stock FROM products WHERE id=$1", item.ProductID).Scan(&productID, &productName, &price, &stock)
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product id %d not found", item.ProductID)
		}

		if err != nil {
			return nil, err
		}
		subtotal := item.Quantity * price
		totalAmount += subtotal

		_, err = tx.Exec("UPDATE products SET stock = stock - $1 WHERE id = $2", item.Quantity, productID)
		if err != nil {
			return nil, err
		}
		details = append(details, model.TransactionDetail{
			ProductID:   productID,
			ProductName: productName,
			Quantity:    item.Quantity,
			Subtotal:    subtotal,
		})
	}

	var transactionID int
	err = tx.QueryRow("INSERT INTO transactions (total_amount) VALUES ($1) RETURNING ID", totalAmount).Scan(&transactionID)
	if err != nil {
		return nil, err
	}

	for i, detail := range details {
		details[i].TransactionID = transactionID
		_, err := tx.Exec("INSERT INTO transaction_details (transaction_id, product_id, quantity, subtotal) VALUES ($1, $2,$3,$4)", transactionID, detail.ProductID, detail.Quantity, detail.Subtotal)
		if err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	res = &model.Transaction{
		ID:          transactionID,
		TotalAmount: totalAmount,
		Details:     details,
	}

	return res, nil
}
