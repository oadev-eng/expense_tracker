package repository

import (
	"tracker/database"
	"tracker/models"
)


type TransactionRepository interface {
	CreateTransaction(transaction *models.Transaction) error
	DuplicateCheck(transaction *models.Transaction) bool
	GetTransactionbyID(id uint) (models.Transaction, error)
	EditTransaction(transaction *models.Transaction) error
	TransactionExist(id uint) bool
	DeleteTransaction(id uint) error
	GetTransactionsbyUserID(id uint) ([]models.Transaction, error)

}

type TransactionRepo struct {}

func (r *TransactionRepo) CreateTransaction(transaction *models.Transaction) error {
	err := database.DB.Create(transaction).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *TransactionRepo) DuplicateCheck(transaction *models.Transaction) bool {
	var count int64
	database.DB.Model(&models.Transaction{}).Where(&models.Transaction{Type: transaction.Type, Category: transaction.Category}).Count(&count)
	return count > 0
}

func (r *TransactionRepo) GetTransactionbyID(id uint) (models.Transaction, error) {
	var transaction models.Transaction
	err := database.DB.Where("id = ?", id).Find(&transaction).Error
	if err != nil {
		return models.Transaction{}, err
	}

	return transaction, nil
}

func (r *TransactionRepo) EditTransaction(transaction *models.Transaction) error {
	err := database.DB.Save(transaction).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *TransactionRepo) TransactionExist(id uint) bool {
	var count int64
	database.DB.Model(&models.Transaction{}).Where("id = ", id).Count(&count)
	return count > 0
}

func (r *TransactionRepo) DeleteTransaction(id uint) error {
	err := database.DB.Delete(&models.Transaction{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *TransactionRepo) GetTransactionsbyUserID(id uint) ([]models.Transaction, error) {
	var trans []models.Transaction
	err := database.DB.Where("user_id = ?", id).Find(&trans).Error
	if err != nil {
		return []models.Transaction{}, err
	}

	return trans, nil
}