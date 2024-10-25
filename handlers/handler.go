package handlers

import (
	"encoding/json"
	"net/http"

	"movie-distribution/models"
)

type Handler struct {
	distributors map[string]*models.Distributor
}

func NewHandler(distributors map[string]*models.Distributor) *Handler {
	return &Handler{
		distributors: distributors,
	}
}

type CheckPermissionRequest struct {
	DistributorName string `json:"distributorName"`
	Region          string `json:"region"`
}

type CheckPermissionResponse struct {
	HasPermission bool `json:"hasPermission"`
}

func (h *Handler) CheckPermission(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CheckPermissionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	distributor, exists := h.distributors[req.DistributorName]
	if !exists {
		http.Error(w, "Distributor not found", http.StatusNotFound)
		return
	}

	result := CheckPermissionResponse{
		HasPermission: distributor.HasPermission(req.Region),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
