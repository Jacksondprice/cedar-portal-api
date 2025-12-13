package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Jacksondprice/cedar-portal-api/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/bookings", handlers.ListBookings)
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8080"
	}
	log.Fatal(http.ListenAndServe(addr, mux))
}
