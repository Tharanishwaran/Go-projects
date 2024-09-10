package main

import (
	"fmt"
	"log"
	"net/http"
)

// Define a custom handler struct
type APIHand struct{}

// Implement the ServeHTTP method on the APIHandler struct
func (APIHand) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request:", r.Method, r.URL.Path)
	// Write "hello world" to the response
	fmt.Fprint(w, "hello universe hi bro hi iron man")
	fmt.Fprint(w, "hello everyone")
	fmt.Fprint(w, "hi bro")
	fmt.Fprint(w, " hi iron man captain america")
	fmt.Fprint(w, " Update ironman captain america thor spiderman")

}

func main() {
	// Start the HTTP server on port 8000 with APIHandler as the handler
	log.Println("Starting server on :8000")
	err := http.ListenAndServe(":8000", &APIHand{})
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
