package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Aleksi! Your url is: %s\n", r.URL.Path)
		//w.WriteHeader(http.StatusNotFound)
	})

	http.ListenAndServe(":8000", nil)
}
