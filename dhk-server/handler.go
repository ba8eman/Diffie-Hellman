package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GET /publicKey
// Generates a p g pair and returns to the caller
func (s *Server) handlePublicKey(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Aaro key chodhikunu ")
	var dh dhKey
	dh.intiPrimes()
	jsonDh, err := json.Marshal(dh)
	if err != nil {
		fmt.Println(err)
	}
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
	err = json.Unmarshal(bodyBytes, &dh)
	fmt.Println(dh)
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
