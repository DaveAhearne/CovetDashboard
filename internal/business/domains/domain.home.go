package domains

import (
	"context"
)

type HomeDomain struct {
}

type HomeUsecase interface {
	Home(ctx context.Context, req HomeDomain) (statusCode int, res HomeDomain)
}
