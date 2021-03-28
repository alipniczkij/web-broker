package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	broker "github.com/alipniczkij/web-broker"
)

func (h *Handler) PutHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handle PUT request")
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	putReq := &broker.PutValue{
		Key:   string(r.URL.Path[1:]),
		Value: r.Form.Get("v"),
	}

	err := h.repo.Queue.Put(putReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprintf(w, "OK")
}

func (h *Handler) GetHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handle GET request")
	timeout, err := strconv.Atoi(r.URL.Query().Get("timeout"))
	if err != nil {
		log.Printf("Timeout error %s", err)
		timeout = 10
	}

	getReq := &broker.GetValue{
		Key: string(r.URL.Path[1:]),
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(timeout)*time.Second)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			http.Error(w, "404 not found.", http.StatusNotFound)
			return
		default:
			v, err := h.repo.Queue.Get(getReq)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			if v == "" {
				continue
			}
			fmt.Fprintf(w, "%v", v)
			return
		}
	}
}
