package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alipniczkij/web-broker/internal/handler"
	"github.com/alipniczkij/web-broker/internal/repository"
)

func main() {
	port := flag.String("port", "8000", "Port")
	storageName := flag.String("storage", "storage.json", "Storage for data")
	log.Printf("Get port %s and storage %s", *port, *storageName)

	repo := repository.NewRepository(*storageName)
	handler := handler.NewHandler(repo)

	log.Print("Start web service")

	serv := http.Server{
		Addr:    net.JoinHostPort("", *port),
		Handler: handler.InitRoutes(),
	}

	go serv.ListenAndServe()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	<-interrupt

	log.Print("Stopping app...")

	timeout, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	err := serv.Shutdown(timeout)
	if err != nil {
		log.Printf("Error when shutdown app: %v", err)
	}

	log.Print("The app stopped")

}

func gracefulShutdown(cancelFunc context.CancelFunc) {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGILL)

	go func() {
		sig := <-sigs
		log.Printf("catched signal: %v", sig)
		cancelFunc()
	}()
}
