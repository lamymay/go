package main

import (
	"log"
	"net/http"
)

//https://jdc.io/3/golang-web-server-hello-world
//https://studygolang.com
//https://blog.csdn.net/asd1126163471/article/details/100987958
func main() {
	mux := http.NewServeMux()

	server := &http.Server{Addr: ":8080", Handler: mux}
	log.Print("111")
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	log.Fatal(server.ListenAndServe())
}
