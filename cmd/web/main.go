package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting server...")
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
}
