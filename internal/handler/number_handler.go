package handler

import (
	"encoding/json"
	"go-crud-evo/internal/service"
	"net/http"
)

type NumberRequest struct {
	Number int `json:"number"`
}

type NumberHandler struct {
	service service.NumberService
}

func NewNumberHandler(service service.NumberService) *NumberHandler {
	return &NumberHandler{service: service}
}

func (h *NumberHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req NumberRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	numbers, err := h.service.ProcessNumber(req.Number)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(numbers)
}
