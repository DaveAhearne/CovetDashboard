package requests

import "covet.digital/dashboard/internal/business/domains"

type HomepageRequest struct {
}

func (hr HomepageRequest) ToDomain() *domains.HomeDomain {
	return &domains.HomeDomain{}
}
