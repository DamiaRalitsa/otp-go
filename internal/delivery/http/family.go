package http

import (
	"bookingtogo/internal/domain"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type FamilyHandler struct {
	FamilyUsecase domain.FamilyUsecase
}

func NewFamilyHandler(familyUsecase domain.FamilyUsecase) *FamilyHandler {
	return &FamilyHandler{
		FamilyUsecase: familyUsecase,
	}
}

func (h *FamilyHandler) GetAllByCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cstID, err := strconv.Atoi(mux.Vars(r)["cst_id"])
	if err != nil {
		http.Error(w, "Invalid customer ID format", http.StatusBadRequest)
		return
	}

	data, err := h.FamilyUsecase.GetAllByCustomer(cstID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(data)
}

func (h *FamilyHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ftID, err := strconv.Atoi(mux.Vars(r)["cst_id"])
	if err != nil {
		http.Error(w, "Invalid family ID format", http.StatusBadRequest)
		return
	}

	data, err := h.FamilyUsecase.GetByID(ftID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(data)
}

func (h *FamilyHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cstID, err := strconv.Atoi(mux.Vars(r)["cst_id"])
	if err != nil {
		http.Error(w, "Invalid customer ID format", http.StatusBadRequest)
		return
	}

	var f domain.Family
	if err := json.NewDecoder(r.Body).Decode(&f); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	f.CustomerID = cstID
	if err := h.FamilyUsecase.Create(&f); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Family member created successfully"})
}

func (h *FamilyHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ftID, err := strconv.Atoi(mux.Vars(r)["ft_id"])
	if err != nil {
		http.Error(w, "Invalid family ID format", http.StatusBadRequest)
		return
	}

	var f domain.Family
	if err := json.NewDecoder(r.Body).Decode(&f); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	f.ID = ftID
	if err := h.FamilyUsecase.Update(&f); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Family member updated successfully"})
}

func (h *FamilyHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ftID, err := strconv.Atoi(mux.Vars(r)["ft_id"])
	if err != nil {
		http.Error(w, "Invalid family ID format", http.StatusBadRequest)
		return
	}

	if err := h.FamilyUsecase.Delete(ftID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Family member deleted successfully"})
}
