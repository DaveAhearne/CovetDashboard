package server

import (
	"context"
	"covet.digital/dashboard/internal/config"
	"covet.digital/dashboard/internal/datasources/drivers"
	"covet.digital/dashboard/internal/http/routes"
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

	// Database connection
	connectionString := drivers.GetDbConnString(
		conf.DatabaseUsername,
		conf.DatabasePassword,
		conf.DatabaseHost,
		conf.DatabasePort,
		conf.DatabaseName)

	connPool, err := drivers.SetupPostgresConnection(connectionString)
	if err != nil {
		return nil, err
	}

	routes.AddWSRoute(mux, connPool, conf).Setup()

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
