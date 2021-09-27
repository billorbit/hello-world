package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello() string {
	return "hello world"
}

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "hello world")
	})

	log.Fatalln(http.ListenAndServe(":8080", nil))
}
