package handlers

import (
	"encoding/json"
	"kasir-api/models"
	"kasir-api/services"
	"net/http"
	"strconv"
)

type CategoryHandler struct {
	service *services.CategoryService
}

func NewCategoryHandler(service *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func (h *CategoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	categories, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(categories)
}

func (h *CategoryHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	category, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(category)
}

func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.Create(&category); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(category)
}

func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.Update(id, &category); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(category)
}

func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Category deleted successfully",
	})
}
