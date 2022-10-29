package main

import (
	"fmt"
	"log"
	"net/http"
)

func enableCors(w *http.ResponseWriter) {
	// (*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Origin", "https://andersonprante.github.io")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	enableCors(&w)
	fmt.Fprintf(w, "Hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static")) // New code
	http.Handle("/", fileServer)                        // New code
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 3000\n")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}

}
