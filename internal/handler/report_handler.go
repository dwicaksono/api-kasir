package handler

import (
	"encoding/json"
	"kasir-api/internal/domain"
	"net/http"
)

type ReportHandler struct {
	usecase domain.ReportUsecase
}

func NewReportHandler(u domain.ReportUsecase) *ReportHandler {
	return &ReportHandler{usecase: u}
}

func (h *ReportHandler) GetReport(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	if startDate == "" || endDate == "" {
		http.Error(w, "start_date and end_date are required", http.StatusBadRequest)
		return
	}

	report, err := h.usecase.GetReport(r.Context(), startDate, endDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(report)
}
