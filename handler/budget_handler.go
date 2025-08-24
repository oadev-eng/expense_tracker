package handler

import (
	"encoding/json"
	"net/http"
	"tracker/middleware"
	"tracker/models"
	"tracker/service"
)

type BudgetHandler struct {
	Service *service.BudgetService
}
func (h *BudgetHandler) CreateBudget(w http.ResponseWriter, r *http.Request) {
	var budget models.Budget
	err := json.NewDecoder(r.Body).Decode(&budget)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	userID, err := middleware.GetUserIDFromToken(r)
	if err != nil {
		http.Error(w, "could not get user_id", http.StatusInternalServerError)
		return
	}

	budget.UserID = userID

	err = h.Service.CreateBudget(&budget)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(budget)

}