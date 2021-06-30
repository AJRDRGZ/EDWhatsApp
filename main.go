package main

import (
	"log"
	"net/http"
)

func main() {
	serveMux := http.NewServeMux()

	// web page
	serveMux.Handle("/", http.FileServer(http.Dir("public")))

	// server websocket
	serveMux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {

	})

	log.Println("running server :8080")
	log.Println(http.ListenAndServe(":8080", serveMux))
}
