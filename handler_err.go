package main

import "net/http"

func handleErr(w http.ResponseWriter, r *http.Request) {
	respondWithErr(w, 400, "Something went Wrong")
}
