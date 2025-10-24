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
	s.repo.GetOrdersByDate(ctx, date)

	s.repo.UpsertDailyReport(ctx, nil)

	return nil, nil
}
