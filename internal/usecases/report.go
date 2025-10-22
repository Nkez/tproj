package usecases

import (
	"context"
	"fmt"
	"time"
	"tproj/internal/models"
	"tproj/internal/repositories"
)

type ReportService interface {
	GenerateDailyReport(ctx context.Context, date time.Time) (*models.DailyReport, error)
}

type reportService struct {
	repo *repositories.ReportRepository
}

func NewReportService(repo *repositories.ReportRepository) ReportService {
	return &reportService{repo: repo}
}

func (s *reportService) GenerateDailyReport(ctx context.Context, date time.Time) (*models.DailyReport, error) {
	fmt.Println("Generate daily report for date")
	orders, err := s.repo.GetOrdersByDate(ctx, date)
	if err != nil {
		return nil, err
	}

	fmt.Println("Orders: ", orders)

	if err := s.repo.UpsertDailyReport(ctx, &models.DailyReport{}); err != nil {
		return nil, err
	}

	return &models.DailyReport{}, nil
}
