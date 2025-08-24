package repository

import (
	"tracker/database"
	"tracker/models"
)

type BudgetRepository interface {
	CreateBudget(budget *models.Budget) error
	DuplicateCheck(budget *models.Budget) bool
}

type BudgetRepo struct{}

func (r *BudgetRepo) CreateBudget(budget *models.Budget) error {
	err := database.DB.Create(budget).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *BudgetRepo) DuplicateCheck(budget *models.Budget) bool {
	var count int64
	database.DB.Model(&models.Budget{}).Where(models.Budget{Category: budget.Category}).Count(&count)
	return count > 0

}