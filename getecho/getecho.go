// getecho.go -- simple service to GET another URL and echo back result.
//
// http://127.0.0.1/https/biztos.com/misc

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const Line = "════════════════════════════════════════════════════════════════════════════════"

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// well we want an URL right? but let's not fuck around with POST.
		dest := strings.TrimPrefix(r.URL.Path, "/")
		if len(dest) == 0 {
			http.Error(w, "no destination in path", 400)
			return
		}

		// the first part must be the protocol and it must be http or https
		if len(dest) < 10 {
			http.Error(w, "insufficient destination in path", 400)
			return
		}
		var url string
		if dest[:6] == "https/" {
			url = fmt.Sprintf("https://%s", dest[6:])
		} else if dest[:5] == "http/" {
			url = fmt.Sprintf("http://%s", dest[5:])

		} else {
			http.Error(w, "path must begin with /(http|https)/", 400)
			return
		}

		resp, err := http.Get(url)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		// hostname is nice to have...
		hostname := os.Getenv("HOSTNAME")
		if hostname == "" {
			hostname = "<UNKOWN SERVER HOST>"
		}

		fmt.Fprintln(w, hostname, time.Now())
		fmt.Fprintln(w, "GET", url)
		fmt.Fprintln(w, resp.Status)
		fmt.Fprintln(w, Line)
		for k, vv := range resp.Header {
			for _, v := range vv {
				fmt.Fprintf(w, "%s: %s\n", k, v)
			}

		}
		fmt.Fprintln(w, Line)
		fmt.Fprintln(w, string(body))

	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
