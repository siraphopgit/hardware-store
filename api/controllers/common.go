package controllers

import (
	"fmt"
	"net/http"
)

func buildCreatedResponse(w http.ResponseWriter, location string) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Location", location)
}

func buildLocation(r *http.Request, id uint64) string {
	return fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, id)
}
