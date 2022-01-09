package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os/signal"
	"strconv"
	"syscall"
)

var (
	Version     = "-"
	Commit      = "-"
	BuildTime   = "-"
	port        int
	showVersion bool
)

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	flag.IntVar(&port, "port", 8000, "Server listening port")
	flag.BoolVar(&showVersion, "version", false, "Show version")
	flag.Parse()

	if showVersion {
		fmt.Printf("Version:\t%s\n", Version)
		fmt.Printf("Commit:\t%s\n", Commit)
		fmt.Printf("Build.time:\t%s\n", BuildTime)
		return
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", HealthHandler)
	s := &http.Server{
		Addr:    net.JoinHostPort("", strconv.Itoa(port)),
		Handler: logRequest(mux),
	}

	go func() {
		log.Println("Starting server...")
		if err := s.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()
	log.Println("Stopping server...")
	if err := s.Shutdown(ctx); err != nil {
		log.Println(err.Error())
	}
}
