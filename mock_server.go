package main

import (
	"fmt"
	"net/http"
)

func main() {
	// This creates a simple web server that acts like AWS S3
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Received %s request for %s\n", r.Method, r.URL.Path)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("<?xml version=\"1.0\" encoding=\"UTF-8\"?><ListBucketResult></ListBucketResult>"))
	})

	fmt.Println("Mock S3 Server started on http://localhost:4566")
	http.ListenAndServe(":4566", nil)
}

