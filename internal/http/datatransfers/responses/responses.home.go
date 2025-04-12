package responses

import "covet.digital/dashboard/internal/business/domains"

type HomeResponse struct {
	Host string
	Port string
}

func NewHomeResponse(u domains.HomeDomain) HomeResponse {
	return HomeResponse{
		Host: u.ApplicationHost,
		Port: u.ApplicationPort,
	}
}
