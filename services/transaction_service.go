package services

import (
	"categories-api/model"
	"categories-api/repositories"
)

type TransactionService struct {
	repo *repositories.TransactionRepository
}

func NewTransactionService(repo *repositories.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) CheckoutItem(items []model.CheckoutItem) (*model.Transaction, error) {
	return s.repo.CreateTransaction(items)
}
