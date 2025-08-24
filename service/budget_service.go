package service

import (
	"errors"
	"tracker/models"
	"tracker/repository"
)

type BudgetService struct {
	Repo repository.BudgetRepository
}

func (s *BudgetService) CreateBudget(budget *models.Budget) error {
	ok := s.Repo.DuplicateCheck(budget)
	if ok {
		return errors.New("duplicate budget")
	}

	err := s.Repo.CreateBudget(budget)
	if err != nil {
		return err
	}

	return nil

}
