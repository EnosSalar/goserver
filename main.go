package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server is running on http://localhost:8080")
	fmt.Println("Visit http://localhost:8080/hello")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
