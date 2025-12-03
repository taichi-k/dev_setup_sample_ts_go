package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/taichi-k/dev_setup_sample_ts_go/backend/internal/service"
)

type UserHandler struct {
	svc service.UserWelcomer
}

func NewUserHandler(svc service.UserWelcomer) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) Welcome(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	if err := h.svc.WelcomeUser(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{"status": "ok"})
}
