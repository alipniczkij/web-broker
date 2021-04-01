package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/alipniczkij/web-broker/internal/models"
)

func (h *Handler) PutHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	putReq := &models.PutValue{
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
	timeout, err := strconv.Atoi(r.URL.Query().Get("timeout"))
	if err != nil {
		log.Printf("Timeout error %s", err)
		timeout = 10
	}

	getReq := &models.GetValue{
		Key: string(r.URL.Path[1:]),
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(timeout)*time.Second)
	defer cancel()

	for {
		time.Sleep(100 * time.Millisecond)
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
