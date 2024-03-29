package main

import "net/http"

func handleReady(w http.ResponseWriter, _ *http.Request) {
	okPayload := struct {
		Status string `json:"status"`
	}{Status: "ok"}

	respondWithJSON(w, http.StatusOK, okPayload)
}

func handleError(w http.ResponseWriter, _ *http.Request) {
	respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}
