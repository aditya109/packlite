package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		log.Println("ok")
		fmt.Fprintf(w, "pong")
	})

	http.ListenAndServe(":8090", nil)
}
