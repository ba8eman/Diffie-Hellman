package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// GET /publicKey
// Generates a p g pair and returns to the caller
func (s *Server) handlePublicKey(w http.ResponseWriter, r *http.Request) {
	var dh dhKey
	dh.intiPrimes()

	jsonDh, err := json.Marshal(dh)
	if err != nil {
		log.Printf("failed to encode public key response: %v", err)
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}

	log.Printf("issued new DH public parameters, id=%s", dh.Id)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonDh)
}

// POST /exchange
func (s *Server) handleExchange(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var dh dhKey
	if err := json.Unmarshal(bodyBytes, &dh); err != nil {
		log.Printf("failed to decode exchange request: %v", err)
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	log.Printf("received exchange request, id=%s", dh.Id)
	otherOverTheWire := dh.OverTheWire
	dh.initPrivateKey()
	dh.findSharedKey(otherOverTheWire)

	respBytes, err := json.Marshal(dh.OverTheWire)
	if err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(respBytes)
}
