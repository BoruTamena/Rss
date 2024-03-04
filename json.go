package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithErr(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5xx error : %v\n", msg)
	}

	type errResponse struct {
		Error string `json : error`
	}

	respondWithJson(w, code, errResponse{
		Error: msg,
	})
}

// the function that define the way server should respond to user request
func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {

	data, err := json.Marshal(payload)

	if err != nil {
		log.Printf("faild to marshal payload \n %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-type", "applicatio/json")
	w.WriteHeader(code)
	w.Write(data)

}
