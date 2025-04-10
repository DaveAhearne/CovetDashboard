package ws

import (
	"github.com/gorilla/websocket"
	"net/http"
)

type WSService interface {
	Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error)
}

type wsService struct {
}

func NewWSService() WSService {
	return &wsService{}
}

func (ws *wsService) Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	var websocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	websocketUpgrader.CheckOrigin = func(r *http.Request) bool {
		// TODO: Remove this, we should probably be checking the origin...
		return true
	}

	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
