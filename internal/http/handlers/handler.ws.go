package handlers

import (
	"covet.digital/dashboard/internal/business/domains"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type WSHandler struct {
	LogUseCase domains.LogUseCase
}

func NewWSHandler(logUseCase domains.LogUseCase) WSHandler {
	return WSHandler{
		LogUseCase: logUseCase,
	}
}

func (wsH WSHandler) HandleStreamLogEvents(w http.ResponseWriter, r *http.Request) {
	conn, err := wsH.LogUseCase.UpgradeConnection(w, r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println(err)
		}
	}(conn)

	logstream, err := wsH.LogUseCase.ListenForLogEvents(r.Context())
	if err != nil {
		_ = conn.Close()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for {
		eventLog, _ := <-logstream
		if err := conn.WriteJSON(eventLog); err != nil {
			fmt.Println("An error occurred writing the log back to the socket", err)
		}
	}
}

func (wsH WSHandler) HandleBasic(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
