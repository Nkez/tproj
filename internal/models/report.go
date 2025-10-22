package models

import "time"

type Order struct {
	ID        int64     `db:"id"`
	UserID    int64     `db:"user_id"`
	Amount    float64   `db:"amount"`
	Status    string    `db:"status"`
	CreatedAt time.Time `db:"created_at"`
}

type DailyReport struct {
	ID           int64     `db:"id"`
	ReportDate   time.Time `db:"report_date"`
	TotalOrders  int64     `db:"total_orders"`
	TotalRevenue float64   `db:"total_revenue"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}
