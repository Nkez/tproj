package repositories

import (
	"context"
	"database/sql"
	"time"
	"tproj/internal/models"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (r *ReportRepository) GetOrdersByDate(ctx context.Context, date time.Time) ([]*models.Order, error) {
	return nil, nil
}

func (r *ReportRepository) UpsertDailyReport(ctx context.Context, report *models.DailyReport) error {
	return nil
}
