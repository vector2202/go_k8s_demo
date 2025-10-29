package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello from a kube")
	})
	fmt.Println("Server running port 8080")
	http.ListenAndServe(":9090", nil)
}
