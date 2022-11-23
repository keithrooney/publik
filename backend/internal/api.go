package internal

import (
	"encoding/json"
	"net/http"
)

type Dispatcher struct {
}

func (d Dispatcher) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		echo(w, req)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func echo(w http.ResponseWriter, req *http.Request) {
	address, err := parse(req.RemoteAddr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(address)
}

func NewDispatcher() Dispatcher {
	return Dispatcher{}
}
