package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/joaofnds/foo/config"
	"github.com/joaofnds/foo/repo"
)

func main() {
	log.Println("Starting the application...")

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

	host, _ := os.Hostname()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		names, err := repo.GetAll(db)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		namesConcated := strings.Join(names, ", ")
		fmt.Fprintf(w, "[%s] names: %s", host, namesConcated)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	s := http.Server{Addr: ":80"}
	go func() {
		log.Fatal(s.ListenAndServe())
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	log.Println("Shutdown signal received, exiting...")

	s.Shutdown(context.Background())
}
