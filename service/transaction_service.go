package service

import (
	"errors"
	"tracker/models"
	"tracker/repository"
)


type TransactionService struct {
	Repo  repository.TransactionRepository
}

func (s *TransactionService) CreateTransaction(transaction *models.Transaction) error {
	// check for duplicate post
	ok := s.Repo.DuplicateCheck(transaction)
	if ok {
		return errors.New("duplicate transaction")
	}

	err := s.Repo.CreateTransaction(transaction)
	if err != nil {
		return err
	}

	return nil 
}

func (s *TransactionService) GetTransactionbyID(id uint) (models.Transaction, error) {
	ok := s.Repo.TransactionExist(id)
	if ok {
		return models.Transaction{}, errors.New("transaction not found")
	}

	
	transaction, err := s.Repo.GetTransactionbyID(id)
	if err != nil {
		return models.Transaction{}, err
	}

	return transaction, nil
	
}

func (s *TransactionService) EditTransaction(transaction *models.Transaction) error {
	err := s.Repo.EditTransaction(transaction)
	if err != nil {
		return err
	}

	return nil

}

func (s *TransactionService) DeleteTransaction(id uint) error {
	ok := s.Repo.TransactionExist(id)
    if ok {
		return errors.New("transaction not found")
	}

	err := s.Repo.DeleteTransaction(id)
	if err != nil {
		return err
	}

	return nil

}

func (s *TransactionService) GetTransactionsbyUserID(id uint) ([]models.Transaction, error) {
	ok := s.Repo.TransactionExist(id)
	if ok {
		return []models.Transaction{}, errors.New("transactions not found")
	}

	trans, err := s.Repo.GetTransactionsbyUserID(id)
	if err != nil {
		return []models.Transaction{}, err
	}

	return trans, nil

}
