package repository

import (
	"context"
	"database/sql"
	"kasir-api/internal/domain"
	"time"
)

type reportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) domain.ReportRepository {
	return &reportRepository{db: db}
}

func (r *reportRepository) GetReport(ctx context.Context, startDate, endDate time.Time) (domain.Report, error) {
	var report domain.Report

	// 1. Total Revenue
	queryRevenue := `
		SELECT COALESCE(SUM(total_amount), 0)
		FROM transactions
		WHERE created_at >= $1 AND created_at <= $2
	`
	err := r.db.QueryRowContext(ctx, queryRevenue, startDate, endDate).Scan(&report.TotalRevenue)
	if err != nil {
		return domain.Report{}, err
	}

	// 2. Total Transactions
	queryCount := `
		SELECT COUNT(*)
		FROM transactions
		WHERE created_at >= $1 AND created_at <= $2
	`
	err = r.db.QueryRowContext(ctx, queryCount, startDate, endDate).Scan(&report.TotalTransactions)
	if err != nil {
		return domain.Report{}, err
	}

	// 3. Best Selling Product
	queryBestSeller := `
		SELECT p.name, COALESCE(SUM(td.quantity), 0) as total_sold
		FROM transaction_details td
		JOIN products p ON td.product_id = p.id
		JOIN transactions t ON td.transaction_id = t.id
		WHERE t.created_at >= $1 AND t.created_at <= $2
		GROUP BY p.name
		ORDER BY total_sold DESC
		LIMIT 1
	`
	err = r.db.QueryRowContext(ctx, queryBestSeller, startDate, endDate).Scan(&report.BestSellingProduct.Name, &report.BestSellingProduct.QuantitySold)
	if err != nil {
		if err == sql.ErrNoRows {
			// No sales in period, return empty best seller
			report.BestSellingProduct = domain.BestSellingProduct{}
		} else {
			return domain.Report{}, err
		}
	}

	return report, nil
}
