package main

import (
	"tproj/internal/interfaces/rest"
	"tproj/internal/interfaces/rest/handlers"
	"tproj/internal/repositories"
	"tproj/internal/usecases"
)

func main() {
	reportRepository := repositories.NewReportRepository(nil)
	reportUsecase := usecases.NewReportService(reportRepository)
	reportHandler := handlers.NewReportHandler(reportUsecase)
	server := rest.NewServer(reportHandler)
	server.Start()
}
