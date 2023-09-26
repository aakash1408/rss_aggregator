package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Fatal("Couldnt MarshaL")
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")

	w.Write(dat)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Respondind with error : - ", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errResponse{
		Error: msg,
	})

}
