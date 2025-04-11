package responses

import "covet.digital/dashboard/internal/business/domains"

type HomeResponse struct {
}

func NewHomeResponse(u domains.HomeDomain) HomeResponse {
	return HomeResponse{}
}
