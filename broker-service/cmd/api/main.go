package main

import (
	"fmt"
	"io"
	"log"
	"log/slog"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Config struct {
	Rabbit *amqp.Connection
}

const webPort = "80"

func main() {
	rabbitConn, err := connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer rabbitConn.Close()

	app := Config{
		Rabbit: rabbitConn,
	}

	slog.Info("Starting broker service on", "webPort:", webPort)

	// define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// start the server
	err = srv.ListenAndServe()
	if err != nil {
		slog.Error("Error while starting server ", "err: ", err)
		os.Exit(0)
	}
}

func connect() (*amqp.Connection, error) {
	var counts int64
	var backoff = 1 * time.Second
	var connection *amqp.Connection

	//do not continue until rabbitmq is ready
	for { //amqp://username:password@localhost
		//c, err := amqp.Dial("amqp://guest:guest@localhost")
		c, err := amqp.Dial("amqp://guest:guest@rabbitmq")
		if err != nil {
			fmt.Println("RabbitMQ not yet ready...")
			counts++
		} else {
			log.Println("Connected to RabbitMQ!")
			connection = c
			break
		}

		if counts > 5 {
			fmt.Println(err)
			return nil, err
		}

		backoff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("backing off")
		time.Sleep(backoff)
		continue
	}

	return connection, nil
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
