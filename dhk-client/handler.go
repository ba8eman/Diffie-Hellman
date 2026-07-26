package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
)

// GET /publicKey
// Generates a p g pair and returns to the caller
func requestPublicKey() (dhKey, error) {
	// 1. Added missing colon to http://
	resp, err := http.Get("http://localhost:8080/publicKey")
	if err != nil {
		log.Printf("Ayyo Error: Failed to connect to server: %v", err)
		return dhKey{}, err // STOP execution here so it doesn't panic below
	}
	// 2. Always close the body when err is nil
	defer resp.Body.Close()

	// 3. Read the actual data sent back by the server
	bodyBytes, _ := io.ReadAll(resp.Body)

	var dh dhKey
	err = json.Unmarshal(bodyBytes, &dh)
	fmt.Println(dh)
	return dh, nil
}

// POST /exchange
func intiateExchange(dh dhKey) *big.Int {
	fmt.Println("Before exchange from client", dh)
	dhBytes, err := json.Marshal(dh)
	if err != nil {
		log.Fatalf("Error modeling struct to JSON: %v", err)
	}

	resp, err := http.Post("http://localhost:8080/exchange", "application/json", bytes.NewBuffer(dhBytes))

	if err != nil {
		log.Fatalf("Failed to send POST request: %v", err)
	}

	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)

	var otherPrivateValue *big.Int
	err = json.Unmarshal(bodyBytes, &otherPrivateValue)
	fmt.Println(otherPrivateValue)
	return otherPrivateValue

}
