package http

import (
	"fmt"
	"net/http"
)

// IndexHandler represents the root / uri of the server
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello - Welcome to PTT")
}
