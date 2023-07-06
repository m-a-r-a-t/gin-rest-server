package app

import (
	"fmt"
	"gin-rest-server/internal/config"
	"gin-rest-server/internal/controller"
	"gin-rest-server/internal/infrastructure/solver"
	"gin-rest-server/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run() {
	s.httpServer.ListenAndServe()

}

func SetupServer(cfg *config.Config, log *slog.Logger) *Server {
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	engine := gin.New()

	engine.Static("/tasks", cfg.StaticFilesPath)

	apiGroup := engine.Group("/api")

	s := Server{&http.Server{
		Addr:         addr,
		Handler:      engine,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	}}

	taskSolverLib := solver.New()

	taskUseCase := usecase.NewTaskUseCase(taskSolverLib)

	taskController := controller.NewTaskController(log, taskUseCase)

	taskController.Register(apiGroup)

	return &s
}
