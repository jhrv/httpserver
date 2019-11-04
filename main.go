package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	bindAddr        string
	defaultResponse string
)

func init() {
	flag.StringVar(&bindAddr, "bind-address", "127.0.0.1:8080", "ip:port where http requests are served")
	flag.StringVar(&defaultResponse, "default-response", "yes\n", "what to respond when recpinged")
	flag.Parse()
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, defaultResponse)
	})

	fmt.Println("running @", bindAddr)
	fmt.Println((&http.Server{Addr: bindAddr}).ListenAndServe())
}
