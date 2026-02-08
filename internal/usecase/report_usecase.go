package usecase

import (
	"context"
	"kasir-api/internal/domain"
	"time"
)

type reportUsecase struct {
	reportRepo domain.ReportRepository
	timeout    time.Duration
}

func NewReportUsecase(r domain.ReportRepository, timeout time.Duration) domain.ReportUsecase {
	return &reportUsecase{
		reportRepo: r,
		timeout:    timeout,
	}
}

func (u *reportUsecase) GetReport(ctx context.Context, startDateStr, endDateStr string) (domain.Report, error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeout)
	defer cancel()

	// Parse dates
	// Layout "2006-01-02" is the reference date for YYYY-MM-DD
	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		return domain.Report{}, err
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		return domain.Report{}, err
	}

	// Adjust endDate to include the entire day (up to 23:59:59)
	// Or query logic can handle it.
	// Let's set endDate to 23:59:59 of that day for better UX if user expects inclusive range
	endDate = endDate.Add(23*time.Hour + 59*time.Minute + 59*time.Second)

	return u.reportRepo.GetReport(ctx, startDate, endDate)
}
