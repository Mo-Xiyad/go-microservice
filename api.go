package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type ApiServer struct {
	service Service
}

func NewApiServer(service Service) *ApiServer {
	return &ApiServer{service: service}
}
func (server *ApiServer) Start(listenAddr string) error {
	http.HandleFunc("/", server.handleGetCatFact)
	// http.HandleFunc("/qr-code", server.handleGetQRCode("https://catfact.ninja/fact"))
	return http.ListenAndServe(listenAddr, nil)
 
}
func (s *ApiServer) handleGetCatFact(w http.ResponseWriter, r *http.Request) {
	fact, err := s.service.GetCatFact(context.Background())
	if err != nil {
		writeJSON(w, map[string]string{"error": err.Error()})
		return
	}
	json.NewEncoder(w).Encode(fact)
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
