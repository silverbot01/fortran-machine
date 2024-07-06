package main

import (
	"log"
	"net/http"

	"github.com/yookoala/gofast"
)

func main() {
	address := "127.0.0.1:9000"
	connFactory := gofast.SimpleConnFactory("tcp", address)

	http.Handle("/", gofast.NewHandler(
		gofast.NewFileEndpoint("")(gofast.BasicSession),
		gofast.SimpleClientFactory(connFactory),
	))

	http.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
