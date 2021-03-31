package controller

import "net/http"

func Check(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
