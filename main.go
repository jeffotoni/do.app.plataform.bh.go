package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/api/v1/ping", Ping)
	println("Server Run v0.0.0 port:8080")
	http.ListenAndServe(":8080", nil)
}

func Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}
