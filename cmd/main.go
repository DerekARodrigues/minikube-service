package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/kardianos/service"
)

type program struct {
	exit chan struct{}
}

var logger service.Logger

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	})

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Starting HTTP server on :8080")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func (p *program) Stop(s service.Service) error {
	close(p.exit)
	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        "GoHttpService",
		DisplayName: "Go HTTP Service",
		Description: "This is an example Go HTTP service using kardianos/service.",
	}

	prg := &program{
		exit: make(chan struct{}),
	}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	// Setup logger for the service
	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}

	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}
