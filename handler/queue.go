package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/alipniczkij/web-broker/utils"
)

type queueValue string

var m sync.Mutex

func (h *Handler) PutHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handle PUT request")
	keyNeeded := string(r.URL.Path[1:])

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	v := r.Form.Get("v")

	m.Lock()
	defer m.Unlock()
	datas, err := utils.ReadJSON(h.storage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if _, found := datas[keyNeeded]; found {
		log.Printf("Try to append value %v", v)
		datas[keyNeeded] = append(datas[keyNeeded], v)
	} else {
		log.Printf("Try to set value %s", []string{v})
		datas[keyNeeded] = []string{v}
	}
	utils.WriteJSON(h.storage, datas)
}

func (h *Handler) GetHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handle GET request")
	keyNeeded := string(r.URL.Path[1:])
	timeout, err := strconv.Atoi(r.URL.Query().Get("timeout"))
	if err != nil {
		log.Printf("Timeout error %s", err)
		timeout = 15
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(timeout)*time.Second)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			http.Error(w, "404 not found.", http.StatusNotFound)
			return
		default:
			v, err := getValue(keyNeeded, h.storage)
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

func getValue(key string, fileName string) (string, error) {
	m.Lock()
	defer m.Unlock()
	datas, err := utils.ReadJSON(fileName)
	if err != nil {
		return "", err
	}
	if value, found := datas[key]; found {
		if len(value) != 0 {
			datas[key] = datas[key][1:]
			err := utils.WriteJSON(fileName, datas)
			if err != nil {
				return "", err
			}
			return value[0], nil
		}
	}
	return "", nil
}
