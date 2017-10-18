package http

import (
	"fmt"
	"log"
	"net/http"
)

// Server is the main http entry point
func Server(port string) {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
