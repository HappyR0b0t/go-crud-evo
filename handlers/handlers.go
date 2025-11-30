package handlers

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
)

type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{db: db}
}

type NumberRequest struct {
	Number int `json:"number"`
}

func (h *Handler) HelloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world!")
}

func (h *Handler) NumberHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "only post method allowed", http.StatusMethodNotAllowed)
		return
	}

	var req NumberRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := h.db.Exec("INSERT INTO numbers (value) VALUES ($1)", req.Number)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rows, err := h.db.Query("SELECT value FROM numbers ORDER BY value ASC")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var numbers []int
	for rows.Next() {
		var n int
		if err := rows.Scan(&n); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		numbers = append(numbers, n)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(numbers)
}
