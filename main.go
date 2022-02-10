package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
	log.WithFields(log.Fields{}).Info("Hello handler invoked!")
}

func main() {
	http.HandleFunc("/hello", hello)

	http.ListenAndServe(":8090", nil)
}
