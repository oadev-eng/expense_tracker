package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tracker/middleware"
	"tracker/models"
	"tracker/service"
)


type TransactionHandler struct {
	Service *service.TransactionService
}

func (h *TransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction models.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	userID, err := middleware.GetUserIDFromToken(r)
	if err != nil {
		http.Error(w, "could not get user id", http.StatusInternalServerError)
		return
	}

	transaction.UserID = userID

	err = h.Service.CreateTransaction(&transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(transaction)

}

func (h *TransactionHandler) GetTransactionbyID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "transaction id is required", http.StatusBadRequest)
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "invalid transaction ID", http.StatusBadRequest)
		return
	}

	transaction, err := h.Service.GetTransactionbyID(uint(idInt))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
    
	
	json.NewEncoder(w).Encode(transaction)

}

func (h *TransactionHandler) EditTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction models.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "transaction ID is required", http.StatusBadRequest)
		return
	}

	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid transaction id", 400)
		return
	}

	transaction.ID = uint(idInt)

	id, err := middleware.GetUserIDFromToken(r)
	if err != nil {
		http.Error(w, "could not get user id", 500)
		return
	}

	transaction.UserID = id

	err = h.Service.EditTransaction(&transaction)
	if err != nil {
		http.Error(w, "could not edit post", 500)
		return
	}

    w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transaction)

}

func (h *TransactionHandler) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "transaction id required", 400)
		return
	}

	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid transaction id", 400)

	}

	err = h.Service.DeleteTransaction(uint(idInt))
	if err != nil {
		http.Error(w, "could not delete tansaction", 500)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func (h *TransactionHandler) GetTransactionsbyUserID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("user_id")
	if id == "" {
		http.Error(w, "user id is required", http.StatusBadRequest)
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	trans, err := h.Service.GetTransactionsbyUserID(uint(idInt))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(trans)
}
