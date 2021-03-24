package handler

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/alipniczkij/web-broker/utils"
)

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
	datas := utils.ReadJSON(h.storage)

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
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	keyNeeded := string(r.URL.Path[1:])

	m.Lock()
	defer m.Unlock()
	datas := utils.ReadJSON(h.storage)

	if value, found := datas[keyNeeded]; found {
		if len(value) != 0 {
			datas[keyNeeded] = datas[keyNeeded][1:]
			utils.WriteJSON(h.storage, datas)
			fmt.Fprintf(w, "%v", value[0])
			return
		}
	}
	http.Error(w, "404 not found.", http.StatusNotFound)
}
