package controller

import (
	"fmt"
	"net/http"
)

func Check(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Health check endpoint\n")
}
