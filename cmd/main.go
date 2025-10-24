package main

import (
	"tproj/internal/interfaces/rest"
	"tproj/internal/interfaces/rest/handlers"
	"tproj/internal/repositories"
	"tproj/internal/usecases"
)

// 1. Получить все ордера за указанный день.
// 2. Посчитать:
// общее количество ордеров;
// общую сумму выручки (order.amount).
// 3. Сформировать структуру DailyReport и сохранить в таблицу daily_reports.
// 4. Если отчет за этот день уже существует — обновить.
// 5. Вернуть итоговую структуру отчета.

func main() {
	reportRepository := repositories.NewReportRepository(nil)
	reportUsecase := usecases.NewReportService(reportRepository)
	reportHandler := handlers.NewReportHandler(reportUsecase)
	server := rest.NewServer(reportHandler)
	server.Start()
}
