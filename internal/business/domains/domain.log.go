package domains

import (
	"context"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

type LogDomain struct {
	Id        string
	AccountId string
	Category  string
	Data      any
	CreatedAt time.Time
}

type LogUseCase interface {
	UpgradeConnection(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error)
	ListenForLogEvents(ctx context.Context) (<-chan LogDomain, error)
}

type LogRepository interface {
	Listen(ctx context.Context) (<-chan LogDomain, error)
}
