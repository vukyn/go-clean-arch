package handler

import (
	"boilerplate-clean-arch/application/domains/auth"
	"boilerplate-clean-arch/config"
)

type authHandlers struct {
	cfg    *config.Config
	authUC auth.UseCase
}

func NewAuthHandlers(cfg *config.Config, authUC auth.UseCase) auth.Handlers {
	return &authHandlers{
		cfg:    cfg,
		authUC: authUC,
	}
}