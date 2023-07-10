package init

import (
	"go-clean-arch/config"
	handler "go-clean-arch/internal/auth/delivery/http"
	"go-clean-arch/internal/auth/repository"
	"go-clean-arch/internal/auth/usecase"

	"github.com/redis/go-redis/v9"
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
	redisClient *redis.Client,
) *Init {
	repo := repository.NewRepo(db)
	redisRepo := repository.NewRedisRepo(redisClient)
	usecase := usecase.NewUseCase(cfg, repo, redisRepo)
	handler := handler.NewHandler(cfg, usecase)
	return &Init{
		Repository: repo,
		Usecase:    usecase,
		Handler:    handler,
	}
}
