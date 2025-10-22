package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"tproj/internal/usecases"
)

type ReportHandler struct {
	report usecases.ReportService
}

func NewReportHandler(report usecases.ReportService) *ReportHandler {
	return &ReportHandler{
		report: report,
	}
}

func (r ReportHandler) GenerateReport(w http.ResponseWriter, req *http.Request) {
	dateStr := req.URL.Query().Get("date")
	if dateStr == "" {
		http.Error(w, "Missing 'date' query parameter (format: 2006-01-02)", http.StatusBadRequest)
		return
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid date format: %v", err), http.StatusBadRequest)
		return
	}

	ctx := req.Context()
	report, err := r.report.GenerateDailyReport(ctx, date)
	if err != nil {
		log.Printf("Failed to generate report: %v", err)
		http.Error(w, fmt.Sprintf("Failed to generate report: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(report); err != nil {
		log.Printf("Failed to encode report: %v", err)
	}
}
