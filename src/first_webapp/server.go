package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	_, err := fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
	if err != nil {
		fmt.Println("Failed to Fprintf")
	}
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Failed to ListenAndServe")
	}
}
