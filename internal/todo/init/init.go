package init

import (
	"boilerplate-clean-arch/config"
	"boilerplate-clean-arch/internal/middleware"
	handler "boilerplate-clean-arch/internal/todo/delivery/http"
	"boilerplate-clean-arch/internal/todo/repository"
	"boilerplate-clean-arch/internal/todo/usecase"

	"gorm.io/gorm"
)

type Init struct {
	Repository repository.IRepository
	Usecase    usecase.IUseCase
	Handler    handler.Handler
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
