package routes

import (
	"covet.digital/dashboard/internal/business/usecases"
	"covet.digital/dashboard/internal/config"
	"covet.digital/dashboard/internal/http/handlers"
	"covet.digital/dashboard/pkg/template"
	"net/http"
)

type homeRoutes struct {
	middleware func(http.HandlerFunc) http.Handler
	handler    handlers.HomeHandler
	mux        *http.ServeMux
}

func AddHomeRoute(mux *http.ServeMux, templateService template.TemplateService, conf config.Config) *homeRoutes {
	homeUsecase := usecases.NewHomeUsecase(conf)
	homeHandler := handlers.NewHomeHandler(homeUsecase, templateService)

	middleware := func(h http.HandlerFunc) http.Handler {
		return http.HandlerFunc(h)
	}

	return &homeRoutes{mux: mux, handler: homeHandler, middleware: middleware}
}

func (r *homeRoutes) Setup() {
	r.mux.Handle("GET /{$}", r.middleware(r.handler.Home))
	r.mux.Handle("GET /", r.middleware(r.handler.RedirectHome))
}
