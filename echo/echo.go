// echo.go -- simple HTTP service to echo request
//
// TODO: fancy decoding of JSON and such stuff on POST, nice to echo that too.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// hostname is nice to have...
		hostname := os.Getenv("HOSTNAME")
		if hostname == "" {
			hostname = "<UNKOWN SERVER HOST>"
		}

		fmt.Fprintln(w, hostname, time.Now(), "--", r.URL.Path)

	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
