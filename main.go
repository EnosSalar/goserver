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

	fmt.Println("Server is running on http://localhost:8082")
	fmt.Println("Visit http://localhost:8082/hello")

	log.Fatal(http.ListenAndServe(":8082", nil))
}
