package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	hostname        string
	defaultResponse string
	port            int
)

func init() {
	flag.StringVar(&hostname, "h", "127.0.0.1", "hostname used to serve requests")
	flag.IntVar(&port, "p", 8080, "port used to serve requests")
	flag.StringVar(&defaultResponse, "default-response", "yes\n", "what to respond when recpinged")
	flag.Parse()
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, defaultResponse)
	})

	bindAddr := fmt.Sprintf("%s:%d", hostname, port)
	fmt.Println("running @", bindAddr)
	fmt.Println((&http.Server{Addr: bindAddr}).ListenAndServe())
}
