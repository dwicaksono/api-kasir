package domain

import (
	"context"
	"time"
)

type Report struct {
	TotalRevenue       float64            `json:"total_revenue"`
	TotalTransactions  int                `json:"total_transactions"`
	BestSellingProduct BestSellingProduct `json:"best_selling_product"`
}

type BestSellingProduct struct {
	Name         string `json:"name"`
	QuantitySold int    `json:"quantity_sold"`
}

type ReportRepository interface {
	GetReport(ctx context.Context, startDate, endDate time.Time) (Report, error)
}

type ReportUsecase interface {
	GetReport(ctx context.Context, startDateStr, endDateStr string) (Report, error)
}
