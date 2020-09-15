package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"github.com/ahrtr/pprof"
)

type myHandler struct {
}

var enableDebug bool // whether enable debug mode or not

func init() {
	flag.BoolVar(&enableDebug, "debug", false, "Whether enable debug mode or not")
}

func (h *myHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	resp.WriteHeader(200)

	response := "hello world!"
	encoder := json.NewEncoder(resp)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(response)
	if err != nil {
		log.Printf("Failed to write response body, %v\n", err)
	}
}

func main() {
	flag.Parse()
	h := &myHandler{}

	// httpmux is the original ServerMux
	httpmux := http.NewServeMux()
	httpmux.Handle("/v2/foo", h)

	// if debug mode is enabled, then register pprof handlers;
	// pprof.RegisterPprof returns a new ServerMux, which wraps the original one
	if enableDebug {
		httpmux = pprof.RegisterPprof(httpmux)
	}

	s := &http.Server{
		Addr:    ":8080",
		Handler: httpmux,
	}

	log.Println("Server is starting...")
	log.Fatal(s.ListenAndServe())
}
