package main

import (
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
)

type Config struct{}

const webPort = "80"

func main() {
	app := Config{}

	slog.Info("Starting broker service on", "webPort:", webPort)

	// define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// start the server
	err := srv.ListenAndServe()
	if err != nil {
		slog.Error("Error while starting server ", "err: ", err)
		os.Exit(0)
	}
}

func init() {
	handlerOpts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	logDir := "/app/logs"
	logFile := "log.txt"
	logfile := filepath.Join(logDir, logFile)
	fmt.Println("logfile ", logfile)
	slog.Info("logpath ", "path: ", logFile)
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		log.Fatalf("Unable to create log directory: %v", err)
	}
	file, err := os.OpenFile(logfile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Println("err ", err)
		log.Fatalf("Unable to find logpath %v", err)
	}
	logger := slog.New(slog.NewJSONHandler(io.Writer(file), handlerOpts))
	slog.SetDefault(logger)
}
