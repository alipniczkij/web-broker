package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/alipniczkij/web-broker/pkg/handler"
)

func main() {
	port := flag.String("port", "8000", "Port")
	storageName := flag.String("storage", "storage.json", "Storage for data")
	log.Printf("Get port %s and storage %s", *port, *storageName)
	handler := handler.NewHandler(*storageName)
	log.Print("Start web service")
	if err := http.ListenAndServe(":"+*port, handler.InitRoutes()); err != nil {
		log.Fatal(err)
	}

}
