package domains

import (
	"context"
)

type HomeDomain struct {
	ApplicationHost string
	ApplicationPort string
}

type HomeUsecase interface {
	Home(ctx context.Context, req HomeDomain) (statusCode int, res HomeDomain)
}
