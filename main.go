package main

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", IndexHandler)

	err := http.ListenAndServe(":9000", mux)
	if err != nil {
		panic(err)
	}
}
