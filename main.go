package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/joaofnds/foo/config"
	"github.com/joaofnds/foo/logger"
	"github.com/joaofnds/foo/repo"
	"github.com/joaofnds/foo/tracing"
	"github.com/opentracing/opentracing-go"
)

func main() {
	logger.InfoLogger().Println("starting the application...")

	ctx := context.Background()
	tracer, closer := tracing.InitTracer("asdf")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	err := config.Parse()
	if err != nil {
		panic(err)
	}

	dbHost := config.GetString("postgres.host")
	dbPort := config.GetString("postgres.port")
	dbuser := config.GetString("postgres.username")
	dbPwd := config.GetString("postgres.password")
	dbName := config.GetString("postgres.database")

	db, err := GetConn(dbHost, dbPort, dbuser, dbPwd, dbName)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	http.HandleFunc("/", rootHandler(ctx, db))
	http.HandleFunc("/health", healthHandler(ctx))

	s := http.Server{Addr: ":80"}
	go func() {
		logger.InfoLogger().Println("starting the server")
		logger.ErrorLogger().Fatal(s.ListenAndServe())
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	logger.InfoLogger().Println("shutdown signal received, exiting...")

	s.Shutdown(context.Background())
}

func rootHandler(ctx context.Context, db *sql.DB) func(http.ResponseWriter, *http.Request) {
	host, _ := os.Hostname()

	return func(w http.ResponseWriter, r *http.Request) {
		span, newCtx := opentracing.StartSpanFromContext(ctx, "rootHandler")
		defer span.Finish()

		names, err := repo.GetAll(newCtx, db)
		if err != nil {
			logger.ErrorLogger().Printf("failed to get things from the database: %+v\n", err)
			w.WriteHeader(500)
			return
		}

		namesConcated := strings.Join(names, ", ")
		fmt.Fprintf(w, "[%s] ids: %s", host, namesConcated)
	}
}

func healthHandler(ctx context.Context) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}
}
