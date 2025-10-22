package rest

import (
	"net/http"
	"time"
	"tproj/internal/interfaces/rest/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	server        *http.Server
	mux           chi.Mux
	reportHandler *handlers.ReportHandler
}

func NewServer(reportHandler *handlers.ReportHandler) *Server {
	r := chi.NewRouter()

	r.Route("/api", func(api chi.Router) {
		api.Use(middleware.Timeout(30 * time.Second))

		api.Route("/reports", func(reports chi.Router) {
			reports.Post("/generate", reportHandler.GenerateReport)
		})
	})

	serv := &Server{
		server: &http.Server{
			Addr:    "localhost:8787",
			Handler: r,
		},
		mux: *r,
	}

	return serv
}

func (r *Server) Start() {
	r.server.ListenAndServe()
}
