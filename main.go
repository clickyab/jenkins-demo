package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s", r.URL)

		fmt.Fprintf(w, "%+v", os.Environ())
	})

	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
