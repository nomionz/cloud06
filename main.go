package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	host, err := os.Hostname()
	if err != nil {
		log.Fatal("failed to get hostname:", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Nikita from %s!", host)
	})

	if err = http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("failed to start server:", err)
	}
}
