package usecases

import (
	"context"
	"covet.digital/dashboard/internal/business/domains"
	"covet.digital/dashboard/pkg/ws"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

type logUsecase struct {
	LogRepository    domains.LogRepository
	WebSocketService ws.WSService
}

func NewLogUsecase(logRepository domains.LogRepository, websocketService ws.WSService) domains.LogUseCase {
	return &logUsecase{
		LogRepository:    logRepository,
		WebSocketService: websocketService,
	}
}

func (l logUsecase) UpgradeConnection(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	return l.WebSocketService.Upgrade(w, r)
}

func (l logUsecase) ListenForLogEvents(ctx context.Context) (<-chan domains.LogDomain, error) {
	return l.LogRepository.Listen(ctx)
}

func (l logUsecase) GetLastWeeksEvents(ctx context.Context) ([]domains.LogDomain, error) {
	return l.LogRepository.GetEventsAfter(ctx, time.Now().AddDate(0, 0, -7))
}
