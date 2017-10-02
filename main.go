package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Set delay value in milliseconds

const delayValue time.Duration = 5000

func main() {

	r := mux.NewRouter()

	r.Path("/fixed-delay").Methods("POST").HandlerFunc(fixedDelayHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Starting middleware webserver")
	fmt.Println(srv.ListenAndServe())
}

func fixedDelayHandler(res http.ResponseWriter, req *http.Request) {

	reqBody, err := ioutil.ReadAll(req.Body)

	if err != nil {
		writeErrorResponse(res, err.Error(), http.StatusInternalServerError)
	}

	fmt.Printf("Delaying response by %s\n", delayValue*time.Millisecond)
	time.Sleep(delayValue * time.Millisecond)
	res.Write(reqBody)
}

type errorView struct {
	Error string `json:"error"`
}

func writeErrorResponse(res http.ResponseWriter, message string, code int) {

	errorView := &errorView{Error: message}
	errorBytes, err := json.Marshal(errorView)

	if err != nil {
		res.WriteHeader(500)
		return
	}

	res.WriteHeader(code)
	writeResponse(res, errorBytes)
}

func writeResponse(res http.ResponseWriter, bytes []byte) {

	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.Write(bytes)
}
