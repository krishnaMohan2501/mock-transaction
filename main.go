package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("POST /upi/pay", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[MOCK-TXN] payment received | X-Risk-Score=%s", r.Header.Get("X-Risk-Score"))
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":        "payment_accepted",
			"transactionId": "TXN-MOCK-001",
		})
	})
	log.Println("[MOCK-TXN] Starting on :9001")
	log.Fatal(http.ListenAndServe(":9001", nil))
}
