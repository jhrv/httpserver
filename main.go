package main

import (
	"flag"
	"fmt"
	"net/http"
)

const BindHost = "127.0.0.1"

var (
	port            int
	defaultResponse string
)

func init() {
	flag.IntVar(&port, "p", 8080, "port where http requests are served")
	flag.StringVar(&defaultResponse, "default-response", "", "what to respond when receiving requests")
	flag.Parse()
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, defaultResponse)
	})

	bindAddr := fmt.Sprintf("%s:%d", BindHost, port)
	fmt.Println("running @", bindAddr)
	fmt.Println((&http.Server{Addr: bindAddr}).ListenAndServe())
}
