package server

import (
	"context"
	"covet.digital/dashboard/internal/config"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	HttpServer *http.Server
}

func NewApp() (*App, error) {
	// initialize config
	conf := config.NewConfig()

	// set up the routes and middleware
	mux := http.NewServeMux()

	server := &http.Server{
		Addr:           fmt.Sprintf("%s:%s", conf.ApplicationHost, conf.ApplicationPort),
		Handler:        mux,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &App{
		HttpServer: server,
	}, nil
}

func (a *App) Run() error {
	go func() {
		log.Printf("success to listen and serve on %s\n", a.HttpServer.Addr)

		if err := a.HttpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Failed to listen and serve: %+v\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// make blocking channel and waiting for a signal
	<-quit
	log.Printf("shutdown server ...\n")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()

	if err := a.HttpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("error when shutdown server: %v", err)
	}

	// catching ctx.Done(). timeout of 5 seconds.
	<-ctx.Done()
	log.Printf("timeout of 5 seconds.")
	log.Printf("server exiting")

	return nil
}

//var upgrader = websocket.Upgrader{
//	ReadBufferSize:  1024,
//	WriteBufferSize: 1024,
//}
//
//var connPool *pgx.Conn
//
//func initDB() {
//	var err error
//	connPool, err = pgx.Connect(context.Background(), "postgres://db_usr:mysecretpassword@0.0.0.0:5432/app")
//	if err != nil {
//		log.Fatalf("Unable to connect to database: %v\n", err)
//	}
//	log.Println("Connected to the database successfully!")
//}
//
//func wsHandler(w http.ResponseWriter, r *http.Request) {
//	upgrader.CheckOrigin = func(r *http.Request) bool {
//		return true
//	}
//
//	conn, err := upgrader.Upgrade(w, r, nil)
//	if err != nil {
//		log.Println("Upgrade error:", err)
//		return
//	}
//	defer conn.Close()
//
//	go func() {
//		for {
//			messageType, p, err := conn.ReadMessage()
//			if err != nil {
//				log.Println("ReadMessage error:", err)
//				return
//			}
//			if err := conn.WriteMessage(messageType, p); err != nil {
//				log.Println("WriteMessage error:", err)
//				return
//			}
//		}
//	}()
//
//	log.Println("Listening for database notifications...")
//	if _, err := connPool.Exec(context.Background(), "LISTEN log_event"); err != nil {
//		log.Fatalf("Failed to listen for notifications: %v", err)
//	}
//
//	for {
//		notification, err := connPool.WaitForNotification(context.Background())
//		if err != nil {
//			log.Println("WaitForNotification error:", err)
//			continue
//		}
//		log.Printf("Received notification: %v", notification.Payload)
//		if err := conn.WriteMessage(websocket.TextMessage, []byte(notification.Payload)); err != nil {
//			log.Println("WriteMessage error:", err)
//			return
//		}
//	}
//}
//
//func main() {
//	initDB()
//	defer connPool.Close(context.Background())
//
//	http.HandleFunc("/ws", wsHandler)
//	log.Fatal(http.ListenAndServe(":1234", nil))
//}
