package application

import (
	"boilerplate-clean-arch/config"
	"database/sql"
	"fmt"

	"gorm.io/gorm"

	authRepo "boilerplate-clean-arch/application/domains/auth/repository"
	authUseCase "boilerplate-clean-arch/application/domains/auth/usecase"
	authHandler "boilerplate-clean-arch/application/domains/auth/delivery/handler"
)

// Server struct
type Server struct {
	echo        *echo.Echo
	cfg         *config.Config
}

// constructor
func NewServer(cfg *config.Config) *Server {
	return &Server{echo: echo.New(), cfg: cfg}
}

func (s *Server) Run() error {

	server := &http.Server{
		Addr:           s.cfg.Server.Port,
		ReadTimeout:    time.Second * s.cfg.Server.ReadTimeout,
		WriteTimeout:   time.Second * s.cfg.Server.WriteTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	go func() {
		s.logger.Infof("Server is listening on PORT: %s", s.cfg.Server.Port)
		if err := s.echo.StartServer(server); err != nil {
			s.logger.Fatalf("Error starting Server: ", err)
		}
	}()

	go func() {
		s.logger.Infof("Starting Debug Server on PORT: %s", s.cfg.Server.PprofPort)
		if err := http.ListenAndServe(s.cfg.Server.PprofPort, http.DefaultServeMux); err != nil {
			s.logger.Errorf("Error PPROF ListenAndServe: %s", err)
		}
	}()

	if err := s.MapHandlers(s.echo); err != nil {
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
	defer shutdown()

	s.logger.Info("Server Exited Properly")
	return s.echo.Server.Shutdown(ctx)
}