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
	ws, err := wsH.LogUseCase.UpgradeConnection(w, r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	logs, err := wsH.LogUseCase.GetLastWeeksEvents(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, log := range logs {
		if err := ws.WriteJSON(log); err != nil {
			fmt.Println("An error occurred writing the log back to the socket", err)
		}
	}

	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println(err)
		}
	}(ws)

	logStream, err := wsH.LogUseCase.ListenForLogEvents(r.Context())
	if err != nil {
		_ = ws.Close()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for {
		eventLog, _ := <-logStream
		if err := ws.WriteJSON(eventLog); err != nil {
			fmt.Println("An error occurred writing the log back to the socket", err)
		}
	}
}

func (wsH WSHandler) HandleBasic(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
