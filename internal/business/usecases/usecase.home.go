package usecases

import (
	"context"
	"covet.digital/dashboard/internal/business/domains"
	"covet.digital/dashboard/internal/config"
	"net/http"
)

type homeUsecase struct {
	config config.Config
}

func NewHomeUsecase(config config.Config) domains.HomeUsecase {
	return &homeUsecase{
		config: config,
	}
}

func (homeUC *homeUsecase) Home(ctx context.Context, req domains.HomeDomain) (statusCode int, res domains.HomeDomain) {
	return http.StatusOK, domains.HomeDomain{
		ApplicationHost: homeUC.config.ExternalAddress,
		ApplicationPort: homeUC.config.ExternalPort,
	}
}
