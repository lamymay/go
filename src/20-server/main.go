package main

import (
	"log"
	"net/http"
)

//https://jdc.io/3/golang-web-server-hello-world
//https://studygolang.com
func main() {
	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	log.Fatal(server.ListenAndServe())
}
