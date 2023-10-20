package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	// everything under code 500 is a client side error, is not interesting for us in the backend
	// but everything with 500+ code is server side error, so important for us
	// so then we need to print it out in the console, so we can fix it
	// client errors should only show the error in the client side - irrelevant for us
	if code > 499 {
		log.Println("Responding with 5XX error: ", msg)
	}

	// In Go, it is usuall when working with JSONs to use a struct and to use a JSON REFLECT TAG
	// which then tells us, how we want our JSON function convert this struct into an JSON Object
	// so this will look something like this then: { "error": "something"}
	// so we define the key, the value is then this what is stored in the Error string
	type errResponse struct {
		Error string `json:"error"`
	}

	// instead of the payload, we will pass the error response to the responseWithJSON function
	respondWithJSON(w, code, errResponse{
		Error: msg,
	})
}

// signature: ResponseWriter from standard http handler, a statuscode in int and an payload as interface
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload) // it will embed whatever we give them as payload and return it as bytes, so we can write it in binary format directly to the http response
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload) // if it failes, we write a status 500 to the header so we now there was a internal errer
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json") // tell the header we are responding with a JSON file
	w.WriteHeader(code)
	w.Write(dat)
}
