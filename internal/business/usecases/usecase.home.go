package usecases

import (
	"context"
	"covet.digital/dashboard/internal/business/domains"
	"net/http"
)

type homeUsecase struct {
}

func NewHomeUsecase() domains.HomeUsecase {
	return &homeUsecase{}
}

func (homeUC *homeUsecase) Home(ctx context.Context, req domains.HomeDomain) (statusCode int, res domains.HomeDomain) {
	return http.StatusOK, domains.HomeDomain{}
}
