package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"math/big"
	"net/http"
)

// GET /publicKey
// Generates a p g pair and returns to the caller
func requestPublicKey() (dhKey, error) {
	resp, err := http.Get("http://localhost:8080/publicKey")
	if err != nil {
		log.Printf("failed to connect to server: %v", err)
		return dhKey{}, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("failed to read public key response: %v", err)
		return dhKey{}, err
	}

	var dh dhKey
	if err := json.Unmarshal(bodyBytes, &dh); err != nil {
		log.Printf("failed to decode public key response: %v", err)
		return dhKey{}, err
	}

	log.Printf("received public parameters from server, id=%s", dh.Id)
	return dh, nil
}

// POST /exchange
func intiateExchange(dh dhKey) *big.Int {
	dhBytes, err := json.Marshal(dh)
	if err != nil {
		log.Fatalf("Error modeling struct to JSON: %v", err)
	}

	log.Printf("sending exchange request, id=%s", dh.Id)
	resp, err := http.Post("http://localhost:8080/exchange", "application/json", bytes.NewBuffer(dhBytes))
	if err != nil {
		log.Fatalf("Failed to send POST request: %v", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("failed to read exchange response: %v", err)
	}

	var otherPrivateValue *big.Int
	if err := json.Unmarshal(bodyBytes, &otherPrivateValue); err != nil {
		log.Fatalf("failed to decode exchange response: %v", err)
	}

	log.Printf("received server's public value, id=%s", dh.Id)
	return otherPrivateValue
}
