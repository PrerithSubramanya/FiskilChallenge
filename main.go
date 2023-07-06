package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync/atomic"
)

var seqNum int64

func handler(w http.ResponseWriter, r *http.Request) {
	num := atomic.AddInt64(&seqNum, 1)
	if num < 0 {
		http.Error(w, "Error: sequence number overflow", http.StatusInternalServerError)
		return
	}
	_, err := fmt.Fprint(w, strconv.FormatInt(num, 10))
	if err != nil {
		http.Error(w, "Error: failed to write response", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error: failed to start server: %v", err)
	}
}
