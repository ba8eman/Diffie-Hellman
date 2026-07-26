package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":8080", "address to listen on, e.g. :8081")
	flag.Parse()

	s := NewServer()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /publicKey", s.handlePublicKey)
	mux.HandleFunc("POST /exchange", s.handleExchange)

	log.Println("listening on", *addr)
	log.Fatal(http.ListenAndServe(*addr, mux))
}
