package main

import (
	"context"
	"github.com/mrnguuyen/go_kube/handlers"
	"github.com/mrnguuyen/go_kube/version"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Printf("Starting the service...\ncommit: %s, buildtime: %s, release: %s",
		version.Commit, version.BuildTime, version.Release)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not set.")
	}

	router := handlers.Router(version.BuildTime, version.Commit, version.Release)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	server := &http.Server{
		Addr:              ":" + port,
		Handler:           router,
	}

	go func() {
		log.Fatal(server.ListenAndServe())
	}()
	log.Print("The service is ready to listen and serve")

	killSignal := <-interrupt
	switch killSignal {
	case os.Interrupt:
		log.Print("Got SIGINT...")
	case syscall.SIGTERM:
		log.Print("Got SIGTERM...")
	}

	log.Print("The service is shutting down...")
	server.Shutdown(context.Background())
	log.Print("Done")
}
