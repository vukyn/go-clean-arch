package init

import (
	"go-clean-arch/config"
	"go-clean-arch/internal/middleware"
	handler "go-clean-arch/internal/todo/delivery/http"
	"go-clean-arch/internal/todo/repository"
	"go-clean-arch/internal/todo/usecase"

	"gorm.io/gorm"
)

type Init struct {
	Repository repository.IRepository
	Usecase    usecase.IUseCase
	Handler    handler.IHandler
}

func NewInit(
	db *gorm.DB,
	cfg *config.Config,
	mw *middleware.MiddlewareManager,
) *Init {
	repo := repository.NewRepo(db)
	usecase := usecase.NewUseCase(repo)
	handler := handler.NewHandler(cfg, usecase, mw)
	return &Init{
		Repository: repo,
		Usecase:    usecase,
		Handler:    handler,
	}
}
