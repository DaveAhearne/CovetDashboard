package handlers

import (
	"covet.digital/dashboard/internal/business/domains"
	"covet.digital/dashboard/internal/http/datatransfers/requests"
	"covet.digital/dashboard/internal/http/datatransfers/responses"
	"covet.digital/dashboard/pkg/template"
	"net/http"
)

type HomeHandler struct {
	usecase         domains.HomeUsecase
	templateService template.TemplateService
}

func NewHomeHandler(usecase domains.HomeUsecase, templateService template.TemplateService) HomeHandler {

	return HomeHandler{
		usecase:         usecase,
		templateService: templateService,
	}
}

func (homeH HomeHandler) Home(w http.ResponseWriter, r *http.Request) {
	req := requests.HomepageRequest{}

	status, res := homeH.usecase.Home(r.Context(), *req.ToDomain())
	homepageResponse := responses.NewHomeResponse(res)

	w.WriteHeader(status)

	homeH.templateService.Execute(w, "views/index", homepageResponse)
}
