package routes

import (
	"covet.digital/dashboard/internal/business/usecases"
	"covet.digital/dashboard/internal/config"
	"covet.digital/dashboard/internal/datasources/repositories"
	"covet.digital/dashboard/internal/http/handlers"
	"covet.digital/dashboard/pkg/ws"
	"github.com/jackc/pgx/v5"
	"net/http"
)

type WsRoutes struct {
	middleware wsRoutesMiddlewares
	handler    handlers.WSHandler
	mux        *http.ServeMux
}

type wsRoutesMiddlewares struct {
}

func AddWSRoute(mux *http.ServeMux, db *pgx.Conn, conf config.Config) *WsRoutes {
	logRepository := repositories.NewLogRepository(db)
	websocketService := ws.NewWSService()

	logUsecase := usecases.NewLogUsecase(logRepository, websocketService)
	logHandler := handlers.NewWSHandler(logUsecase)

	return &WsRoutes{mux: mux, handler: logHandler, middleware: wsRoutesMiddlewares{}}
}

func (r *WsRoutes) Setup() {
	r.mux.HandleFunc("GET /ws", r.handler.HandleStreamLogEvents)
	r.mux.HandleFunc("GET /test", r.handler.HandleBasic)
}
