package main

import (
	"fmt"
	"log"
	"net/http"
)

func firstEndPointHandler(w http.ResponseWriter, r *http.Request) {
	message := "this is the first endpoint"
	_, err := w.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func secondEndPointHandler(w http.ResponseWriter, r *http.Request) {
	message := "second endpoint"
	_, err := w.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func thirdEndPointHandler(w http.ResponseWriter, r *http.Request) {
	message := "second endpoint"
	_, err := w.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello!")
	})
	http.HandleFunc("/first", firstEndPointHandler)
	http.HandleFunc("/second", secondEndPointHandler)
	http.HandleFunc("/third", thirdEndPointHandler)
	fmt.Printf("Starting server at port 8081\n")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
