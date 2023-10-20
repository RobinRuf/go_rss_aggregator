package main

import "net/http"

// in Go you have to use this function signature if you want to define a http handler the way the go standard library expects
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}
